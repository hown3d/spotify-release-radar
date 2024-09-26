package spotify

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zmb3/spotify/v2"
)

func (c Client) CreatePlaylist(ctx context.Context, name string, artists []string) (spotify.ID, error) {
	playlist, err := c.spotify.CreatePlaylistForUser(ctx, c.userID, name, "", true, false)
	if err != nil {
		return "", fmt.Errorf("creating playlist: %w", err)
	}
	err = c.AddNewTracksToPlaylist(ctx, playlist.ID, artists)
	if err != nil {
		return "", err
	}
	return playlist.ID, nil
}

func (c *Client) AddNewTracksToPlaylist(ctx context.Context, playlistID spotify.ID, artists []string) error {
	existingTracks, err := c.getPlaylistTracks(ctx, playlistID)
	if err != nil {
		return err
	}

	var newTracks []spotify.SimpleTrack
	for _, artist := range artists {
		tracks, err := c.getNewTracksOfArtist(ctx, artist)
		if err != nil {
			return fmt.Errorf("getting new tracks of artist %v: %w", artist, err)
		}
		slog.Debug(fmt.Sprintf("found %d tracks of artist %v", len(tracks), artist))
		for _, track := range tracks {
			_, exists := existingTracks[track.ID]
			if !exists {
				newTracks = append(newTracks, track)
			}
		}
	}
	err = c.addTracksToPlaylist(ctx, playlistID, newTracks)
	if err != nil {
		return fmt.Errorf("adding tracks to playlist: %w", err)
	}
	return nil
}

func (c *Client) addTracksToPlaylist(ctx context.Context, playlistID spotify.ID, tracks []spotify.SimpleTrack) error {
	trackIDs := make([]spotify.ID, len(tracks))
	for index, track := range tracks {
		trackIDs[index] = track.ID
	}
	_, err := c.spotify.AddTracksToPlaylist(ctx, playlistID, trackIDs...)
	return err
}

func (c *Client) getPlaylistTracks(ctx context.Context, playlistID spotify.ID) (map[spotify.ID]struct{}, error) {
	tracks := map[spotify.ID]struct{}{}
	page, err := c.spotify.GetPlaylistItems(ctx, playlistID)
	if err != nil {
		return nil, err
	}
	for {
		for _, item := range page.Items {
			// skip podcast episodes
			track := item.Track.Track
			if track == nil {
				continue
			}
			tracks[track.ID] = struct{}{}
		}
		err = c.spotify.NextPage(ctx, page)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return tracks, nil
}

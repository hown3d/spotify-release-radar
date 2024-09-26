package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/zmb3/spotify/v2"
)

type createPlaylistRequest struct {
	Name    string   `json:"name"`
	Artists []string `json:"artists"`
}

type createPlaylistResponse struct {
	ID string `json:"ID"`
}

func createPlaylistHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		client, err := spotifyClient(r)
		if err != nil {
			authErr := apiError{statusCode: http.StatusUnauthorized, message: err.Error()}
			authErr.handle(w)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			reqErr := apiError{
				statusCode: http.StatusBadRequest,
				message:    err.Error(),
			}
			reqErr.handle(w)
			return
		}

		var req createPlaylistRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			reqErr := apiError{
				statusCode: http.StatusBadRequest,
				message:    err.Error(),
			}
			reqErr.handle(w)
			return
		}

		playlistID, err := client.CreatePlaylist(r.Context(), req.Name, req.Artists)
		if err != nil {
			var (
				spotifyErr spotify.Error
				apiErr     apiError
			)
			if errors.Is(err, &spotifyErr) {
				apiErr.message = spotifyErr.Message
				apiErr.statusCode = spotifyErr.Status
			} else {
				apiErr.statusCode = http.StatusInternalServerError
				apiErr.message = err.Error()
			}
			apiErr.handle(w)
			return
		}

		resp := createPlaylistResponse{
			ID: string(playlistID),
		}
		respBody, err := json.Marshal(resp)
		if err != nil {
			apiErr := apiError{
				statusCode: http.StatusInternalServerError,
				message:    err.Error(),
			}
			apiErr.handle(w)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(respBody)
	}
}

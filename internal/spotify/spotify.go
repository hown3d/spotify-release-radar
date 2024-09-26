package spotify

import (
	"context"
	"fmt"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

type Client struct {
	spotify *spotify.Client
	userID  string
}

func NewClient(t *oauth2.Token) (*Client, error) {
	ctx := context.Background()
	httpClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(t))
	c := &Client{
		spotify: spotify.New(httpClient),
	}
	userID, err := c.getUserID(ctx)
	if err != nil {
		return c, fmt.Errorf("getting id of user: %w", err)
	}
	c.userID = userID
	return c, nil
}

func (c *Client) getUserID(ctx context.Context) (string, error) {
	user, err := c.spotify.CurrentUser(ctx)
	if err != nil {
		return "", fmt.Errorf("getting current user: %w", err)
	}
	return user.ID, nil
}

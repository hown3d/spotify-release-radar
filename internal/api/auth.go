package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/hown3d/spotify-release-radar/internal/spotify"
	"golang.org/x/oauth2"
)

func spotifyClient(r *http.Request) (*spotify.Client, error) {
	authHeader, ok := r.Header["Authorization"]
	if !ok {
		return nil, errors.New("missing Authorization header")
	}

	if len(authHeader) < 1 {
		return nil, errors.New("Authorization header missing value")
	}

	bearer := authHeader[0]
	splitBearer := strings.Split(bearer, " ")
	if len(splitBearer) != 2 {
		return nil, errors.New("Authorization header is in invalid format")
	}
	if splitBearer[0] != "Bearer" {
		return nil, errors.New("Authorization header does not use auth type \"Bearer\"")
	}
	token := &oauth2.Token{
		AccessToken: splitBearer[1],
	}

	return spotify.NewClient(token)
}

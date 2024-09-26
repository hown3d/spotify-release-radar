package main

import (
	"log/slog"
	"os"

	"github.com/hown3d/spotify-release-radar/internal/api"
)

func main() {
	addr := ":8080"
	api := api.NewAPI(addr)
	slog.Info("running api", "addr", addr)
	if err := api.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

# Spotify Release Radar

This project aims to provide a self managed release radar for your favorite artists.

# How to run

```sh
go run main.go
```

The API will be accessible on localhost:8080

## Implemented

- Create playlist with the newest tracks of your artists

  ```sh
  curl -H "Authorization: Bearer <TOKEN>" -XPOST localhost:8080/playlist -d '{"name": "myplaylist", "artists": ["linkin park"]}'
  ```

## TODOs

- Store state of playlist in database to update on a regular basis
- Allow updates to playlists with new artists

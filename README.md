# Thumbs generator

Video thumbs generator using Go and ![FFMPEG](https://ffmpeg.org/)

## Dependencies
You should have the FFMPEG and the Go language programming installed in your machine locally.

- FFMPEG - [install here](https://ffmpeg.org/download.html).
- Go - [install here](https://go.dev/dl/)

## Running the application
```
go run main.go
```

### Backlog/Next feature
- [ ] Sync image name with the segment media sequence
- [ ] Serve the image in an HTTP server
- [ ] Control the go routines that run generate HLS Playlist and thumbs with context
package worker

import (
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func Execute() error {
	log.Info().Msg("Running generate playlist with FFMPEG...")

	args := []string{
		"-loglevel", "error",
		"-i", "video/coelhao.mp4",
		"-c", "copy",
		"-b:v", "2M",
		"-maxrate", "4M",
		"-f", "hls",
		"-hls_time", "2",
		"-hls_list_size", "5",
		"-hls_flags", "delete_segments",
		"-hls_segment_filename",
		"output/segment_%03d.ts",
		"output/playlist.m3u8",
	}

	cmd := exec.Command("ffmpeg", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

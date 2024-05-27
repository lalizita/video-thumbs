package cmd

import (
	"github.com/lalizita/thumbs/internal/worker"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func RunWorker() *cobra.Command {
	return &cobra.Command{
		Use:   "live",
		Short: "Run worker to generate ffmpeg live stream",
		Run: func(*cobra.Command, []string) {
			if err := worker.Execute(); err != nil {
				log.Error().Err(err).Msg("Failed to generate playlist")
			}
		},
	}
}

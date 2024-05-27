package cmd

import (
	"github.com/lalizita/thumbs/internal/server"
	"github.com/spf13/cobra"
)

func RunServer() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Run worker to generate ffmpeg live stream",
		Run: func(*cobra.Command, []string) {
			server.Execute()
		},
	}
}

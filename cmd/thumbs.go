package cmd

import (
	"github.com/lalizita/thumbs/internal/thumbs"
	"github.com/spf13/cobra"
)

func RunThumbs() *cobra.Command {
	return &cobra.Command{
		Use:   "thumbs",
		Short: "Run thumbs generator using ffmpeg",
		Run: func(*cobra.Command, []string) {
			thumbs.Execute()
		},
	}
}

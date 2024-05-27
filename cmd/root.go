package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "gothumbs",
		Short:         "Run your thumbs generator",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(
		RunWorker(),
		RunThumbs(),
		RunServer(),
	)

	return rootCmd
}

func Execute() {
	if err := os.MkdirAll("output", os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("Failed to create output path directory")
	}
	if err := os.MkdirAll("thumbs", os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("Failed to create thumbs path directory")
	}

	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

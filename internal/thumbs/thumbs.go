package thumbs

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func Execute() {
	log.Info().Msg("Running generate thumbs with FFMPEG...")

	for {
		files, err := os.ReadDir("output")
		if err != nil {
			fmt.Println("err opening directory")
		}

		if len(files) == 0 {
			fmt.Println("files not found")
		}

		for _, file := range files {
			if !file.IsDir() {
				fileName := file.Name()

				if fileName != "playlist.m3u8" {
					file := fileName[:len(fileName)-3]
					segment := fmt.Sprintf("output/%s", fileName)
					outputImage := fmt.Sprintf("thumbs/%s.png", file)

					_, err := os.Stat(outputImage)
					if err == nil {
						continue
					}

					fmt.Println(file)

					args := []string{
						"-loglevel", "error",
						"-i", segment,
						"-frames:v", "1",
						"-vf", "fps=1",
						outputImage,
					}
					cmd := exec.Command("ffmpeg", args...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr

					if err := cmd.Run(); err != nil {
						fmt.Println("Erro ao criar miniatura:", err)
						continue
					}
				}
			}
		}
	}
}

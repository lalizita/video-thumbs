package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	go generateHLSPlaylist()
	go generateThumbs()

	e := echo.New()
	e.GET("/healthcheck", healthCheck)

	e.Static("/coelho", "output")
	e.Static("/thumbs", "thumbs")

	e.Logger.Fatal(e.Start(":" + "8000"))
}

// ffmpeg -i output/playlist.m3u8 -vf fps=1 thumbs/%f.png
func generateThumbs() {
	log.Info("Running generate thumbs....")
	// Abrir o diretório
	files, err := os.ReadDir("output")
	if err != nil {
		fmt.Println("Erro ao abrir o diretório:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("Nenhum segmento encontrado")
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()

			if fileName != "playlist.m3u8" {
				file := fileName[:len(fileName)-3]
				segment := fmt.Sprintf("output/%s", fileName)
				path := fmt.Sprintf("thumbs/%s.png", file)

				fmt.Println(file)

				args := []string{
					"-loglevel", "error",
					"-i", segment,
					"-frames:v", "1",
					"-vf", "fps=1",
					path,
				}
				cmd := exec.Command("ffmpeg", args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err != nil {
					log.Fatalf("failed generate thumbs", err)
					break
				}
			}
		}
	}

}

// ffmpeg  -stream_loop -1 -i "video/coelhao.mp4" -c copy -f hls -hls_time 5 -hls_list_size 10 -hls_flags 10 -hls_segment_type mpegts -hls_segment_filename "output/segment_%03d.ts" "output/playlist.m3u8"
func generateHLSPlaylist() {
	log.Info("Running generate playlist....")

	args := []string{
		"-loglevel", "error",
		"-stream_loop", "-1",
		"-i", "video/coelhao.mp4",
		"-c", "copy",
		"-f", "hls",
		"-hls_time", "5",
		"-hls_list_size", "10",
		"-hls_flags", "delete_segments",
		"-hls_segment_type", "mpegts",
		"-hls_segment_filename", "output/segment_%03d.ts", "output/playlist.m3u8",
	}
	cmd := exec.Command("ffmpeg", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("failed generate hls", err)
	}
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}

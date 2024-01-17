package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func main() {
	mp4Files, err := filepath.Glob("./input/*.mp4")
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}

	fmt.Println("mp4Files: ", mp4Files)

	for _, inputFile := range mp4Files {
		outputFile := filepath.Join("./output", fmt.Sprintf("%s.gif", filepath.Base(inputFile)))

		promt := exec.Command("ffmpeg", "-i", inputFile, "-vf", "fps=25,scale=680:-1:flags=lanczos", "-c:v", "gif", outputFile)
		// promt :=

		// exec.Command("ffmpeg", "-i", inputFile, "-vf", "fps=15,scale=480:-1:flags=lanczos", "-pix_fmt", "pal8", "-dither", "bayer", "-coalesce", "-b:v", "500k", "-c:v", "gif", outputFile)

		// exec.Command("ffmpeg", "-i", inputFile, "-vf", "fps=30", "-b:v", "2M", "-c:v", "gif", outputFile)

		err := promt.Run()
		if err != nil {
			fmt.Printf("Error converting %s: %v\n", inputFile, err)
		} else {
			fmt.Printf("Conversion of %s completed successfully.\n", inputFile)
		}
	}

}

/*
	prompt

	ffmpeg -ss 30 -t 3 -i input.mp4 \
    -vf "fps=10,scale=320:-1:flags=lanczos,split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse" \
    -loop 0 output.gif
*/

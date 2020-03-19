package main

import (
// 	"fmt"

	"./lib"

	"strconv"
)

func main() {
	down := lib.NewDownload(
		"http://demodownload.image-line.com/flstudio/flstudio_win_20.6.2.1549.exe",
		"./files/flstudio_win_20.6.2.1549.exe",
	)

	go down.Do()

	for prog := range down.Progress {
		/* do something with progress
		   NOTE: this occurs every 
		   single time the file is written to! */
	}
}
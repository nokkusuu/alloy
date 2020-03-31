package main

import (
	"./lib"
)

func main() {
	pack, err := lib.OpenPackFile("./modpack.json")
	if err != nil {
		panic(err)
	}

	dls, err := pack.GetDownloads()
	if err != nil {
		panic(err)
	}

	lib.DownloadMods(dls)
}
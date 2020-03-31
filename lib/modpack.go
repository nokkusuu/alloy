package lib

import (
	"fmt"
	"strconv"
	"strings"
)

func (pack *JSONPack) GetDownloads() ([]Mod, error) {
	var modids []Mod
	
	for _, mod := range pack.Mods {
		if mod.Release == 0 { // if there is no release id specified, get one
			var bestCandidate Mod
			
			mods, err := GetModFiles(mod.CurseName)
			if err != nil {
				return []Mod{}, err
			}
			
			for _, fil := range mods {
				if (fil.ID > bestCandidate.ID) && (fil.MinecraftVersion[0] == pack.MCVersion) {
					bestCandidate = fil
				}
			}

			modids = append(modids, bestCandidate)
		} else {
			var bestCandidate Mod

			mods, err := GetModFiles(mod.CurseName)
			if err != nil {
				return []Mod{}, err
			}
			
			for _, fil := range mods {
				if fil.ID == mod.Release {
					bestCandidate = fil
				}
			}

			modids = append(modids, bestCandidate)
		}
	}

	return modids, nil
}

func DownloadMods(mods []Mod) {
	for _, mod := range mods {
		down := NewDownload(
			GetFileURL(mod.ID, mod.Filename),
			"./files/" + mod.Filename,
		)
	
		go down.Do()
	
		for prog := range down.Progress {
			/* do something with progress
			   NOTE: this occurs every 
			   single time the file is written to! */
			fmt.Println(prog)
		}
	}
}

// TODO: make this a method of Mod
func GetFileURL(id uint, filename string) string {
	var url string = "https://edge.forgecdn.net/files/"
	var sid string = strconv.FormatUint(uint64(id), 10)

	url += sid[:len(sid)-3] + "/" + strings.TrimLeft(sid[len(sid)-3:], "0")
	url += "/" + filename

	return url
}
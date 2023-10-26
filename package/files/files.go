package files

import (
	"os"
	"strings"
)

// Getting files in path

func GetMusicInFolder(path string) ([]string, error) {
	var data []string

	// Getting all mp3 files
	files, err := os.ReadDir(path)
	if err != nil {
		return data, err
	}
	// Getting all mp3 files paths
	for _, file := range files {
		if !file.IsDir() {
			if strings.Split(file.Name(), ".")[1] == "mp3" {
				data = append(data, path+"/"+file.Name())
			}
		}
	}
	return data, nil
}

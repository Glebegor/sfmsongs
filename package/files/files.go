package files

import (
	"fmt"
	"image"
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
	fmt.Println(data)
	return data, nil
}

func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

package files

import (
	"os"
)

// Getting files in path
func GetMusicInFolder() ([]os.DirEntry, error) {
	// Path
	path := "C:/Users/glebe/Music/Music"
	files, err := os.ReadDir(path)
	if err != nil {
		return files, err
	}
	return files, nil
}

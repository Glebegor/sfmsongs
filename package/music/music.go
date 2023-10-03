package music

import "io/fs"

// Function to play music

type Music struct {
	Music_is_play bool
	musicPlayer
}
type musicPlayer interface {
	StartPlayMusic() error
	PlayMusic(filePath string) error
	StopPlayMusic() error
}

func (m *Music) StartPlayMusic(filePath fs.DirEntry) error {
	return nil
}
func (m *Music) PlayMusic(filePath fs.DirEntry) error {
	return nil
}
func (m *Music) StopPlayMusic() error {
	return nil
}

// func PlayMusic(filePath string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Create a beep streamer
// 	streamer, format, err := mp3.Decode(file)
// 	if err != nil {
// 		return err
// 	}
// }

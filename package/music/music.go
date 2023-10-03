package music

import (
	"fmt"
	"sync"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

// Function to play music

type Music struct {
	MusicIsPlay bool
	Player      *oto.Player
	Decoder     *mp3.Decoder
	Paused      bool
	PauseMutex  sync.Mutex
}

func (m *Music) StartPlayMusic(filePath string, sec int) error {
	err := m.PlayMusic(filePath, sec)
	if err != nil {
		return err
	}
	return nil
}

func (m *Music) PlayMusic(filePath string, sec int) error {
	fmt.Print("Music started play\n")
	return nil
}

func (m *Music) PauseMusic() {
	fmt.Print("Music ended play\n")
}

func (m *Music) StopPlayMusic() {

}

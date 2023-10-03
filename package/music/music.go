package music

import (
	"fmt"
	"time"

	oto "github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

// truct to play music
type Music struct {
	IsPlay   bool
	dec      *mp3.Decoder
	context  *oto.Context
	player   *oto.Player
	paused   bool
	position time.Duration
	stopCh   chan struct{} // Channel for signaling stop
}

// NewPlayer
func NewPlayer() *Music {
	return &Music{
		stopCh: make(chan struct{}),
	}
}

// Starting of play
func (m *Music) StartPlayMusic(filePath string, sec int) error {
	go func() {
		err := m.PlayMusic(filePath, sec)
		if err != nil {
			fmt.Printf("Error playing music: %v\n", err)
		}
	}()

	return nil
}

// Playing of music
func (m *Music) PlayMusic(filePath string, sec int) error {
	return nil
}

// Pausing of music
func (m *Music) PauseMusic() {
	fmt.Print("\nMusic paused\n")
}

// Stopping of music
func (m *Music) StopPlayMusic() {
	fmt.Print("\nMusic stopped\n")
}

// Length of music in seconds
func (m *Music) LengthOfMusic(filePath string) int {
	return 0
}

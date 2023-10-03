package music

import (
	"fmt"
	"io"
	"os"
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

func (m *Music) StartPlayMusic(filePath string) error {
	go func() {
		err := m.PlayMusic(filePath)
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}()

	return nil
}
func (m *Music) PlayMusic(filePath string) error {
	fmt.Print("Music on")
	// Open mp3 file
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Decodding of file
	m.Decoder, err = mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	// Creating a new context for playing of music
	c, err := oto.NewContext(m.Decoder.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	// Testing
	m.Player = c.NewPlayer()

	fmt.Printf("Length: %d[bytes]\n", m.Decoder.Length())

	if _, err := io.Copy(m.Player, m.Decoder); err != nil {
		return err
	}
	m.MusicIsPlay = true

	return nil
}

func (m *Music) PausePlayMusic() error {
	m.PauseMutex.Lock()
	defer m.PauseMutex.Unlock()

	if m.Player != nil && !m.Paused {
		fmt.Println("Music paused")
		m.Player.Pause()
		m.Paused = true
	}
}

func (m *Music) StopPlayMusic() {
	fmt.Println("Music off")

	m.PauseMutex.Lock()
	defer m.PauseMutex.Unlock()

	if m.Player != nil {
		m.Player.Close()
	}

	if m.Decoder != nil {
		m.Decoder.Close()
	}

	m.MusicIsPlay = false
	m.Paused = false
}

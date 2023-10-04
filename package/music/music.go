package music

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/widget"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

// truct to play music
type Music struct {
	IsPlay          bool
	dec             *mp3.Decoder
	player          oto.Player
	paused          bool
	StopCh          chan struct{} // Channel for signaling stop
	PositionCh      chan time.Duration
	SecondOfPlaying int
}

// NewPlayer
func NewPlayer() *Music {
	return &Music{
		StopCh:          make(chan struct{}),
		PositionCh:      make(chan time.Duration),
		SecondOfPlaying: 0,
	}
}

// Starting of play
func (m *Music) StartPlayMusic(filePath string, sec int, float1 *widget.Float, w *app.Window) error {
	m.SecondOfPlaying = 0
	m.StopCh = make(chan struct{})
	go func() {
		err := m.PlayMusic(filePath, sec)
		if err != nil {
			fmt.Printf("Error playing music: %v\n", err)
		}
	}()
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				m.SecondOfPlaying++
				float1.Value = float32(m.SecondOfPlaying)
				w.Invalidate()
			case <-m.StopCh:
				return
			}
		}
	}()
	return nil
}

// Playing of music
func (m *Music) PlayMusic(filePath string, sec int) error {
	fmt.Print("\nMusic playing\n")
	m.SecondOfPlaying = sec

	filesBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic("Cant read my bytes file.")
	}

	filesBytesReader := bytes.NewReader(filesBytes)

	m.dec, err = mp3.NewDecoder(filesBytesReader)
	if err != nil {
		panic("Error in decoder: " + err.Error())
	}
	fmt.Print("Decoder worked\n")

	sapmlingRate := 44100
	numOfChannels := 2
	audioBitDepth := 2
	otoCtx, readyChan, err := oto.NewContext(m.dec.SampleRate(), numOfChannels, audioBitDepth)
	if err != nil {
		panic("Oto read failed: " + err.Error())
	}

	<-readyChan

	// Playing of music by second of start
	m.player = otoCtx.NewPlayer(m.dec)
	newPos, err := m.player.(io.Seeker).Seek(int64(sec)*int64(sapmlingRate)*4, io.SeekStart)
	if err != nil {
		panic("player.Seek failed: " + err.Error())
	}
	println("Player is now at position:", newPos)
	m.player.Play()
	// We can wait for the sound to finish playing using something like this
	for m.player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	return nil
}

// Pausing of music
func (m *Music) PauseMusic() {
	fmt.Print("\nMusic paused\n")
	// m.player.Pause()
	// if m.paused == false {
	// 	m.paused = true
	// 	m.player.Pause()
	// } else {
	// 	m.player.Reset()
	// }

}

// Stopping of music
func (m *Music) StopPlayMusic() {
	if m.player != nil {
		m.player.Close()
	}
	if m.dec != nil {
		m.dec = nil
	}
	// Close the channel to signal that no more "stop" signals will be sent
	close(m.StopCh)

}

// Length of music in seconds
func (m *Music) LengthOfMusic(filePath string) (int, error) {
	// Openning file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	// Creating of decoder
	dec, err := mp3.NewDecoder(file)
	if err != nil {
		return 0, err
	}

	m.dec = dec
	// getting a len of music
	samples := dec.Length() / 4
	audioLength := int(samples) / int(dec.SampleRate())

	return int(audioLength), nil
}

func (m *Music) GetSec() int {
	return m.SecondOfPlaying
}

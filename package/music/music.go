package music

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"gioui.org/widget"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

// truct to play music
type Music struct {
	IsPlay bool
	dec    *mp3.Decoder
	player oto.Player
	paused bool
	stopCh chan struct{} // Channel for signaling stop
}

// NewPlayer
func NewPlayer() *Music {
	return &Music{
		stopCh: make(chan struct{}),
	}
}

// Starting of play
func (m *Music) StartPlayMusic(filePath string, sec int, float1 *widget.Float) error {
	go func() {
		err := m.PlayMusic(filePath, sec)
		if err != nil {
			fmt.Printf("Error playing music: %v\n", err)
		}
	}()
	go func() {
		for {
			float1.Value += 1
			time.Sleep(1 * time.Second)
		}
	}()

	return nil
}

// Playing of music
func (m *Music) PlayMusic(filePath string, sec int) error {
	fmt.Print("\nMusic playing\n")

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
	fmt.Print("\nMusic stopped\n")
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

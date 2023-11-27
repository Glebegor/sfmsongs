package music

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"gioui.org/app"
	"gioui.org/widget"
	"github.com/dhowden/tag"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

// truct to play music
type Music struct {
	dec    *mp3.Decoder
	Player oto.Player

	SecondOfPlaying int
	SoundVol        float64
	// Channels
	StopCh     chan struct{} // Channel for signaling stop
	PositionCh chan time.Duration

	// mods
	Repeat       bool
	PlayPlaylist bool
	Paused       bool
	IsPlay       bool
}

type PlayMusic struct {
	timeInSec int
	title     string
	artist    string
	album     string
	genre     string
	pic       *tag.Picture
}

func (p *PlayMusic) GetPicture() []byte {
	return p.pic.Data
}

// NewPlayer
func NewPlayer() *Music {
	// Play and Stop vars of the music
	IsPlay := false
	Repeat := false
	PlayPlaylist := false
	return &Music{
		StopCh:          make(chan struct{}),
		PositionCh:      make(chan time.Duration),
		SecondOfPlaying: 0,
		IsPlay:          IsPlay,
		Repeat:          Repeat,
		PlayPlaylist:    PlayPlaylist,
	}
}

// Starting of play
func (m *Music) StartPlayMusic(filePath string, sec int, secOfEnd int, float1 *widget.Float, w *app.Window) error {
	m.SecondOfPlaying = 0
	m.StopPlayMusic()
	m.StopCh = make(chan struct{})
	go func() {
		err := m.PlayMusic(filePath, sec, secOfEnd, float1, w)
		if err != nil {
			fmt.Printf("Error playing music: %v\n", err)
		}
	}()
	return nil
}

// Playing of music
func (m *Music) PlayMusic(filePath string, sec int, secOfEnd int, float1 *widget.Float, w *app.Window) error {
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
	sapmlingRate := 44100
	numOfChannels := 2
	audioBitDepth := 2
	otoCtx, readyChan, err := oto.NewContext(m.dec.SampleRate(), numOfChannels, audioBitDepth)
	if err != nil {
		panic("Oto read failed: " + err.Error())
	}

	// Sending signal to the thread
	<-readyChan
	// Playing of music by second of start
	m.Player = otoCtx.NewPlayer(m.dec)

	m.Player.SetVolume(float64(m.SoundVol))

	newPos, err := m.Player.(io.Seeker).Seek(int64(sec)*int64(sapmlingRate)*4, io.SeekStart)
	if err != nil {
		panic("player.Seek failed: " + err.Error())
	}
	println("Player is now at position:", newPos)
	m.Player.Play()

	// Creating timer to change seconds
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// We can wait for the sound to finish playing using something like this
	for m.Player.IsPlaying() {
		select {

		case <-ticker.C:
			fmt.Println("Second of playing: ", m.SecondOfPlaying)
			float1.Value = float32(m.SecondOfPlaying)
			w.Invalidate()
			// Check if the specified end second is reached
			if m.SecondOfPlaying == secOfEnd {
				if m.Repeat == true {
					m.StartPlayMusic(filePath, 0, secOfEnd, float1, w)
				}
				m.StopCh <- struct{}{} // Send stop signal
			}
			m.SecondOfPlaying++
		case <-m.StopCh:
			fmt.Println("Stopping playback...")
			return nil
		}

	}
	<-m.StopCh
	return nil
}

// Stopping of music
func (m *Music) StopPlayMusic() {
	if m.Player != nil {
		m.Player.Close()
	}
	if m.dec != nil {
		m.dec = nil
	}
	// Close the channel to signal that no more "stop" signals will be sent
	select {
	case _, ok := <-m.StopCh:
		if !ok {
			return
		} else {
			close(m.StopCh)

		}
	default:
		close(m.StopCh)
	}
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
func (m *Music) SetVolume(soundVol float64) {
	m.Player.SetVolume(float64(soundVol))
}
func (m *Music) GetSec() int {
	return m.SecondOfPlaying
}

// Getting name of file
func (m *Music) GetName(filePath string) string {
	return filepath.Base(filePath)
}

func (m *Music) GetInfoAboutSong(filePath string, lenSec int) (PlayMusic, error) {
	var data PlayMusic

	file, err := os.Open(filePath)
	if err != nil {
		return PlayMusic{}, err
	}
	defer file.Close()

	// Read ID3 tags from the file
	tag, err := tag.ReadFrom(file)
	if err != nil {
		return PlayMusic{}, err
	}

	// Access various tag information
	data.title = tag.Title()
	data.artist = tag.Artist()
	data.album = tag.Album()
	data.genre = tag.Genre()
	data.timeInSec = lenSec
	data.pic = tag.Picture()
	return data, nil
}

# SFMSongs
## About project
<p>SFMSongs(or Smart For Me Songs) application to listening your the lovest music.</p>
<p>In this app using framework gio on language Golang.</p>
<p>Project has 3 main packages, music, files, layouts.</p>
<p>To switch layouts app using switch cases.</p>

## Work with application through makefile
<p>How to install make on windows https://linuxhint.com/install-use-make-windows/</p>
<b>Download dependencies</b>

```
make install-depends
```
<b>Build for windows</b>

```
make build-win
```
<b>Running</b>

```
make run
```

## Architecture of application
<b>App</b> 
```
{
	th              *material.Theme // theme of buttons
	ops             op.Ops // 
	w               *app.Window // Window 
	FolderWithMusic string // Main folder with music
	chosenLayer     string // Chosen layout that you can see
}
```
<b>App</b> 
```
{
	th              *material.Theme // theme of buttons
	ops             op.Ops // 
	w               *app.Window // Window 
	FolderWithMusic string // Main folder with music
	chosenLayer     string // Chosen layout that you can see
}
```

### Layouts

<b>MainLayout</b>
```
{
    // Buttons of layout
	optionButton    widget.Clickable
	optionSongs     widget.Clickable
	optionPlayLists widget.Clickable
	optionThisSong  widget.Clickable

    // Test variable on Option layout
	IsOptionTrue bool
}
```

<b>OptionsLayout</b>
```
{
    // Main layout
	LayoutMain        layout.Dimensions

    // Buttons
	musicFolderInput  widget.Editor
	musicFolderButton widget.Clickable

    // Main folder
	MainFolder        string
}
```

<b>SongsLayout</b>
```
{
	// Buttons
	playPrevButton        widget.Clickable
	playCurrencyButton    widget.Clickable
	playNextButton        widget.Clickable
	sliderLenOfMusic      widget.Float
	sliderSoundVol        widget.Float
	optionsButton         widget.Clickable
	repeatButton          widget.Clickable
	playAllPlaylistButton widget.Clickable

	// Params of music
	idOfMusicInDir int
	lenOfMusic     float32

	// Player
	Player           *music.Music // Player
	pathOfMusic      string // Path to music folder
	MusicThatPlaying music.PlayMusic // Music that Playing right now
	MusicArray       []string // Set of all musics
}
```

<b>SongListLayout</b>
```
{
	PlayLists []SongsListItem
	SongsBTNS []widget.Clickable
}
```
### Help structures

<b>SongsListItem</b>
```
{
	Title       string
	Description string
	Img         string
	MusicPath   string
	LenOfMusic  int
}
```

### How working Player?
#### Structures
<b>Music</b>
```
{
	dec    *mp3.Decoder // Decoder for music
	Player oto.Player // Player struct

	SecondOfPlaying int // Music starting play from this second
	SoundVol        float64 // Sound volume
	// Channels
	StopCh     chan struct{} // Channel for signaling stop
	PositionCh chan time.Duration // Position of th channel

	// mods
	Repeat       bool // Repeat checker
	PlayPlaylist bool // Playlist checker
	Paused       bool // Pause checker
	IsPlay       bool // If that play checker
}
```

<b>PlayMusic</b>
```
{
	timeInSec int // Time in seconds
	title     string 
	artist    string
	album     string
	genre     string
	pic       *tag.Picture
}
```

#### Methods of structures
```
func (p *PlayMusic) GetPicture() []byte // Get picture of music
func NewPlayer() *Music // Initialization of music structure
func (m *Music) StartPlayMusic(filePath string, sec int, secOfEnd int, float1 *widget.Float, w *app.Window) error // Start of playing music
func (m *Music) PlayMusic(filePath string, sec int, secOfEnd int, float1 *widget.Float, w *app.Window) error // Playing of music
func (m *Music) StopPlayMusic() // Stop of playing music
func (m *Music) LengthOfMusic(filePath string) (int, error) // Getting length of music in seconds
func (m *Music) SetVolume(soundVol float64) // Setting volume 
func (m *Music) GetSec() int // Getting second of playing right now
func (m *Music) GetName(filePath string) string // Getting name of music
func (m *Music) GetInfoAboutSong(filePath string, lenSec int) (PlayMusic, error) // Getting full information about music
```
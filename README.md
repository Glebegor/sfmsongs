# SFMSongs
## About project
<p>SFMSongs(or Smart For Me Songs) application to listening your the lovest music.</p>
<p>In this app I using framework gio on language Golang</p>

## Build application
### Building to exe
### Running

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
## 
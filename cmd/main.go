package main

import (
	"fmt"
	"log"
	"os"
	"sfmsonds/package/layouts"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type App struct {
	// App
	th              *material.Theme
	ops             op.Ops
	w               *app.Window
	FolderWithMusic string
	chosenLayer     string
	// Buttons
	// playPrevButton        widget.Clickable
	// playCurrencyButton    widget.Clickable
	// playNextButton        widget.Clickable
	// sliderLenOfMusic      widget.Float
	// sliderSoundVol        widget.Float
	// optionsButton         widget.Clickable
	// repeatButton          widget.Clickable
	// playAllPlaylistButton widget.Clickable

	// states of app

	// Params of music
	// idOfMusicInDir int
	// lenOfMusic     float32

	// // Player
	// Player           *music.Music
	// pathOfMusic      string
	// MusicThatPlaying music.PlayMusic
}

func main() {
	go func() {
		App := new(App)
		// New window
		App.w = app.NewWindow(
			app.Title("SFMSongs"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)
		// Drawwing of window
		if err := App.draw(App.w); err != nil {
			log.Fatal(err)
		}
		// EXIT command
		os.Exit(0)
	}()
	app.Main()
}

func (a *App) draw(w *app.Window) error {
	a.chosenLayer = "songs"
	a.th = material.NewTheme()
	a.FolderWithMusic = "C:/Users/glebe/Music/Music"
	// Initialization of options
	optionLayer := layouts.NewOptionLayout()
	optionLayer.MainFolder = a.FolderWithMusic

	// Another layouts
	mainLayer := new(layouts.MainLayout)
	songsLayer := layouts.NewSongsLayout(a.FolderWithMusic)
	playListLayer := layouts.NewPlayListsLayout()
	songListLayer := layouts.NewSongListLayout()

	// listen for events in the window.
	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&a.ops, e)

			songsLayer.ListenEvents(w)

			a.chosenLayer = mainLayer.ListenEvents(w, a.chosenLayer)
			optionLayer.ListenEvents(songsLayer, w)
			playListLayer.ListenEvents(w)

			// Showing layouts
			switch a.chosenLayer {
			case "songs":
				mainLayer.Layout(gtx, a.th, songsLayer.Init(gtx, a.th))
			case "options":
				mainLayer.Layout(gtx, a.th, optionLayer.Init(gtx, a.th))
			case "playList":
				mainLayer.Layout(gtx, a.th, playListLayer.Init(gtx, a.th))
			case "songList":
				mainLayer.Layout(gtx, a.th, songListLayer.Init(gtx, a.th, songsLayer.MusicArray))
			default:
				fmt.Println(a.chosenLayer)
			}

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}

	}
	return nil
}

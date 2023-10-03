package main

import (
	"fmt"
	"log"
	"os"
	"sfmsonds/package/files"
	"sfmsonds/package/music"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		// New window
		w := app.NewWindow(
			app.Title("SFMSongs"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		// EXIT command
		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window) error {

	// UI variable
	var ops op.Ops

	// Play buttons
	var (
		playPrevButton     widget.Clickable
		playCurrencyButton widget.Clickable
		playNextButton     widget.Clickable
	)

	// Play and Stop vars of the music
	musicPlayer := new(music.Music)
	musicPlayer.Music_is_play = false
	musicArray, _ := files.GetMusicInFolder()

	// Themes
	th := material.NewTheme()

	// listen for events in the window.
	for e := range w.Events() {
		switch e := e.(type) {

		case system.FrameEvent:
			if playCurrencyButton.Clicked() {
				if musicPlayer.Music_is_play {
					fmt.Print("Off music")
					musicPlayer.Music_is_play = false
					if err := musicPlayer.StopPlayMusic(); err != nil {
						fmt.Errorf(err.Error())
					}
				} else {
					fmt.Print("On music")
					musicPlayer.Music_is_play = true
					if err := musicPlayer.StartPlayMusic(musicArray[0]); err != nil {
						fmt.Errorf(err.Error())
					}
				}
			}
			gtx := layout.NewContext(&ops, e)
			// Events

			// Creating layout
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				// Buttons play next, play current, play prev
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.Spacing(layout.Middle),
						}.Layout(gtx,
							// Prev btn
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									margins := layout.Inset{Right: unit.Dp(10), Left: unit.Dp(10)}
									return margins.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											playPBtn := material.Button(th, &playPrevButton, "Prev")
											return playPBtn.Layout(gtx)
										},
									)
								},
							),
							// Play btn
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									margins := layout.Inset{Right: unit.Dp(10), Left: unit.Dp(10)}
									return margins.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											playPBtn := material.Button(th, &playCurrencyButton, "Play")
											return playPBtn.Layout(gtx)
										},
									)
								},
							),
							// Next btn
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									margins := layout.Inset{Right: unit.Dp(10), Left: unit.Dp(10)}
									return margins.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											playPBtn := material.Button(th, &playNextButton, "Next")
											return playPBtn.Layout(gtx)
										},
									)
								},
							),
						)
					},
				), // End of buttons layout
				// Spaces
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				), // End of spaces
			)
			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}

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
		idOfMusicInDir     int
	)
	float1 := new(widget.Float)
	// Play and Stop vars of the music
	musicPlayer := music.NewPlayer()
	musicPlayer.IsPlay = false

	path := "C:/Users/glebe/Music/Music"

	musicArray, err := files.GetMusicInFolder(path)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if len(musicArray) != 0 {
		idOfMusicInDir = 0
	}
	// Themes
	th := material.NewTheme()

	// listen for events in the window.
	for e := range w.Events() {
		switch e := e.(type) {

		case system.FrameEvent:
			if playPrevButton.Clicked() {
				if idOfMusicInDir != 0 {
					idOfMusicInDir -= 1
				}
			}
			if playNextButton.Clicked() {
				if idOfMusicInDir != len(musicArray)-1 {
					idOfMusicInDir += 1
				}
			}
			if playCurrencyButton.Clicked() {

				if musicPlayer.IsPlay {

					musicPlayer.IsPlay = false
					musicPlayer.PauseMusic()
				} else if !musicPlayer.IsPlay {

					musicPlayer.IsPlay = true
					lenMus, _ := musicPlayer.LengthOfMusic(musicArray[idOfMusicInDir])
					fmt.Print(lenMus)
					musicPlayer.StartPlayMusic(musicArray[idOfMusicInDir], int(float1.Value))
				}
			}
			gtx := layout.NewContext(&ops, e)
			// Events

			// Creating layout
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				// Slider to change start sec of music
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.Spacing(layout.Middle),
						}.Layout(gtx,
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
										layout.Flexed(1, material.Slider(th, float1, 0, 180).Layout),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.UniformInset(unit.Dp(18)).Layout(gtx,
												material.Body1(th, fmt.Sprintf("%.0f", float1.Value)).Layout,
											)
										}),
									)
								},
							),
						)
					},
				),
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

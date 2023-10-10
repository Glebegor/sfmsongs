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

type App struct {
	// App
	th  *material.Theme
	ops op.Ops
	w   *app.Window

	// Buttons
	playPrevButton        widget.Clickable
	playCurrencyButton    widget.Clickable
	playNextButton        widget.Clickable
	sliderLenOfMusic      widget.Float
	sliderSoundVol        widget.Float
	optionsButton         widget.Clickable
	repeatButton          widget.Clickable
	playAllPlaylistButton widget.Clickable

	// states of app

	// Params of music
	idOfMusicInDir int
	lenOfMusic     float32

	// Player
	Player           *music.Music
	pathOfMusic      string
	MusicThatPlaying music.PlayMusic
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
	// Play and Stop vars of the music
	a.Player = music.NewPlayer()
	a.Player.IsPlay = false

	a.pathOfMusic = "C:/Users/glebe/Music/Music"

	// Getting all files in dir
	musicArray, err := files.GetMusicInFolder(a.pathOfMusic)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if len(musicArray) != 0 {
		a.idOfMusicInDir = 0
	} else {
		fmt.Errorf("Dont have mp3 files in folder")
	}
	a.Player.Repeat = false
	a.Player.PlayPlaylist = false
	// Themes
	a.th = material.NewTheme()

	// Started music
	lenMus, _ := a.Player.LengthOfMusic(musicArray[a.idOfMusicInDir])
	a.lenOfMusic = float32(lenMus)

	// listen for events in the window.
	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			// Play by second
			if !a.sliderLenOfMusic.Dragging() && a.sliderLenOfMusic.Changed() {
				if a.Player.IsPlay {
					a.Player.StartPlayMusic(musicArray[a.idOfMusicInDir], int(a.sliderLenOfMusic.Value), int(a.lenOfMusic), &a.sliderLenOfMusic, a.w)
				}
				if !a.Player.IsPlay {
					a.Player.IsPlay = false
				}
			}
			// Changing volume
			if a.sliderSoundVol.Dragging() {
				if a.Player.Player != nil {
					a.Player.SoundVol = float64(a.sliderSoundVol.Value)
					a.Player.SetVolume(float64(a.sliderSoundVol.Value))
				}
			}
			// RepeatButton
			if a.repeatButton.Clicked() {
				if a.Player.Repeat {
					a.Player.Repeat = false
				} else if !a.Player.Repeat {
					a.Player.Repeat = true
				}
			}
			// PlayListButton
			if a.repeatButton.Clicked() {
				if a.Player.PlayPlaylist {
					a.Player.PlayPlaylist = false
				} else if !a.Player.PlayPlaylist {
					a.Player.PlayPlaylist = true
				}
			}
			// Prev music
			if a.playPrevButton.Clicked() {
				if a.idOfMusicInDir != 0 {
					if a.Player.IsPlay == false {
						a.idOfMusicInDir -= 1
						fmt.Print(musicArray[a.idOfMusicInDir])
						lenMus, _ := a.Player.LengthOfMusic(musicArray[a.idOfMusicInDir])
						a.lenOfMusic = float32(lenMus)
						a.sliderLenOfMusic.Value = 0

					} else if a.Player.IsPlay == true {
						a.idOfMusicInDir -= 1
						lenMus, err := a.Player.LengthOfMusic(musicArray[a.idOfMusicInDir])
						if err != nil {
							fmt.Errorf(err.Error())
						}
						a.lenOfMusic = float32(lenMus)
						a.sliderLenOfMusic.Value = 0

						a.Player.StartPlayMusic(musicArray[a.idOfMusicInDir], int(a.sliderLenOfMusic.Value), int(a.lenOfMusic), &a.sliderLenOfMusic, a.w)
					}
				}
			}

			// Next music
			if a.playNextButton.Clicked() {
				if a.idOfMusicInDir != len(musicArray) {
					if a.Player.IsPlay == false {
						a.idOfMusicInDir += 1
						fmt.Print(musicArray[a.idOfMusicInDir])
						lenMus, _ := a.Player.LengthOfMusic(musicArray[a.idOfMusicInDir])
						a.lenOfMusic = float32(lenMus)
						a.sliderLenOfMusic.Value = 0
					} else if a.Player.IsPlay == true {
						a.idOfMusicInDir += 1
						fmt.Print(musicArray[a.idOfMusicInDir])
						lenMus, _ := a.Player.LengthOfMusic(musicArray[a.idOfMusicInDir])
						a.lenOfMusic = float32(lenMus)
						a.sliderLenOfMusic.Value = 0

						a.Player.StartPlayMusic(musicArray[a.idOfMusicInDir], int(a.sliderLenOfMusic.Value), int(a.lenOfMusic), &a.sliderLenOfMusic, a.w)
					}
				}
			}

			if a.playCurrencyButton.Clicked() {
				if a.Player.IsPlay {
					a.Player.IsPlay = false
					a.Player.StopPlayMusic()
				} else if !a.Player.IsPlay {
					a.Player.IsPlay = true
					lenMus, _ := a.Player.LengthOfMusic(musicArray[a.idOfMusicInDir])
					a.lenOfMusic = float32(lenMus)

					a.Player.StartPlayMusic(musicArray[a.idOfMusicInDir], int(a.sliderLenOfMusic.Value), lenMus, &a.sliderLenOfMusic, a.w)
				}
			}

			gtx := layout.NewContext(&a.ops, e)
			// Creating layout
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceBetween,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis: layout.Horizontal,
						}.Layout(gtx,
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									optionsBtn := material.Button(a.th, &a.optionsButton, "Options")
									return optionsBtn.Layout(gtx)
								},
							),
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.Spacing(layout.Middle),
						}.Layout(gtx,
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									text := a.Player.GetName(musicArray[a.idOfMusicInDir])
									return material.Body1(a.th, text[:len(text)-4]).Layout(gtx)
								},
							),
						)
					},
				),
				// Sec and maxSec
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.Spacing(layout.Middle),
						}.Layout(gtx,
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.UniformInset(unit.Dp(18)).Layout(gtx,
												material.Body1(a.th, fmt.Sprintf("%.0f/%.0f", a.sliderLenOfMusic.Value, a.lenOfMusic)).Layout,
											)
										}),
									)
								},
							),
						)
					},
				),
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
										layout.Flexed(1, material.Slider(a.th, &a.sliderLenOfMusic, 0, a.lenOfMusic).Layout))
								},
							),
						)
					},
				),
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
										layout.Flexed(1, material.Slider(a.th, &a.sliderSoundVol, 0, 1).Layout))
								},
							),
						)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.Spacing(layout.Middle),
						}.Layout(gtx,
							// Repeat btn
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									margins := layout.Inset{Right: unit.Dp(10), Left: unit.Dp(10)}
									return margins.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											playPBtn := material.Button(a.th, &a.repeatButton, "Repeat")
											return playPBtn.Layout(gtx)
										},
									)
								},
							),
							// All Playlist btn
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									margins := layout.Inset{Right: unit.Dp(10), Left: unit.Dp(10)}
									return margins.Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											playPBtn := material.Button(a.th, &a.playAllPlaylistButton, "Play Playlist")
											return playPBtn.Layout(gtx)
										},
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
											playPBtn := material.Button(a.th, &a.playPrevButton, "Prev")
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
											playPBtn := material.Button(a.th, &a.playCurrencyButton, "Play")
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
											playPBtn := material.Button(a.th, &a.playNextButton, "Next")
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

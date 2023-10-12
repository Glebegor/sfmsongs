package layouts

import (
	"fmt"
	"sfmsonds/package/files"
	"sfmsonds/package/music"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type SongsLayout struct {
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
	Player           *music.Music
	pathOfMusic      string
	MusicThatPlaying music.PlayMusic
	MusicArray       []string
}

// func (a *SongsLayout) ListenEvents(w *app.Window) {
// 	if !a.sliderLenOfMusic.Dragging() && a.sliderLenOfMusic.Changed() {
// 		if a.Player.IsPlay {
// 			a.Player.StartPlayMusic(a.MusicArray[a.idOfMusicInDir], int(a.sliderLenOfMusic.Value), int(a.lenOfMusic), &a.sliderLenOfMusic, w)
// 		}
// 		if !a.Player.IsPlay {
// 			a.Player.IsPlay = false
// 		}
// 	}
// 	// Changing volume
// 	if a.sliderSoundVol.Dragging() {
// 		if a.Player.Player != nil {
// 			a.Player.SoundVol = float64(a.sliderSoundVol.Value)
// 			a.Player.SetVolume(float64(a.sliderSoundVol.Value))
// 		}
// 	}
// 	// RepeatButton
// 	if a.repeatButton.Clicked() {
// 		if a.Player.Repeat {
// 			a.Player.Repeat = false
// 		} else if !a.Player.Repeat {
// 			a.Player.Repeat = true
// 		}
// 	}
// 	// PlayListButton
// 	if a.repeatButton.Clicked() {
// 		if a.Player.PlayPlaylist {
// 			a.Player.PlayPlaylist = false
// 		} else if !a.Player.PlayPlaylist {
// 			a.Player.PlayPlaylist = true
// 		}
// 	}
// 	// Prev music
// 	if a.playPrevButton.Clicked() {
// 		if a.idOfMusicInDir != 0 {
// 			if a.Player.IsPlay == false {
// 				a.idOfMusicInDir -= 1
// 				fmt.Print(a.MusicArray[a.idOfMusicInDir])
// 				lenMus, _ := a.Player.LengthOfMusic(a.MusicArray[a.idOfMusicInDir])
// 				a.lenOfMusic = float32(lenMus)
// 				a.sliderLenOfMusic.Value = 0

// 			} else if a.Player.IsPlay == true {
// 				a.idOfMusicInDir -= 1
// 				lenMus, err := a.Player.LengthOfMusic(a.MusicArray[a.idOfMusicInDir])
// 				if err != nil {
// 					fmt.Errorf(err.Error())
// 				}
// 				a.lenOfMusic = float32(lenMus)
// 				a.sliderLenOfMusic.Value = 0

// 				a.Player.StartPlayMusic(a.MusicArray[a.idOfMusicInDir], int(a.sliderLenOfMusic.Value), int(a.lenOfMusic), &a.sliderLenOfMusic, w)
// 			}
// 		}
// 	}

// 	// Next music
// 	if a.playNextButton.Clicked() {
// 		if a.idOfMusicInDir != len(a.MusicArray) {
// 			if a.Player.IsPlay == false {
// 				a.idOfMusicInDir += 1
// 				fmt.Print(a.MusicArray[a.idOfMusicInDir])
// 				lenMus, _ := a.Player.LengthOfMusic(a.MusicArray[a.idOfMusicInDir])
// 				a.lenOfMusic = float32(lenMus)
// 				a.sliderLenOfMusic.Value = 0
// 			} else if a.Player.IsPlay == true {
// 				a.idOfMusicInDir += 1
// 				fmt.Print(a.MusicArray[a.idOfMusicInDir])
// 				lenMus, _ := a.Player.LengthOfMusic(a.MusicArray[a.idOfMusicInDir])
// 				a.lenOfMusic = float32(lenMus)
// 				a.sliderLenOfMusic.Value = 0

// 				a.Player.StartPlayMusic(a.MusicArray[a.idOfMusicInDir], int(a.sliderLenOfMusic.Value), int(a.lenOfMusic), &a.sliderLenOfMusic, w)
// 			}
// 		}
// 	}

// 	if a.playCurrencyButton.Clicked() {
// 		if a.Player.IsPlay {
// 			a.Player.IsPlay = false
// 			a.Player.StopPlayMusic()
// 		} else if !a.Player.IsPlay {
// 			a.Player.IsPlay = true
// 			lenMus, _ := a.Player.LengthOfMusic(a.MusicArray[a.idOfMusicInDir])
// 			a.lenOfMusic = float32(lenMus)

//				a.Player.StartPlayMusic(a.MusicArray[a.idOfMusicInDir], int(a.sliderLenOfMusic.Value), lenMus, &a.sliderLenOfMusic, w)
//			}
//		}
//	}
func NewSongsLayout(gtx layout.Context, th *material.Theme) *SongsLayout {
	// Buttons
	var playPrevButton widget.Clickable
	var playCurrencyButton widget.Clickable
	var playNextButton widget.Clickable
	var sliderLenOfMusic widget.Float
	var sliderSoundVol widget.Float
	var optionsButton widget.Clickable
	var repeatButton widget.Clickable
	var playAllPlaylistButton widget.Clickable

	// Play and Stop vars of the music
	Player := music.NewPlayer()
	Player.IsPlay = false

	pathOfMusic := "C:/Users/glebe/Music/Music"

	// Getting all files in dir
	musicArray, err := files.GetMusicInFolder(pathOfMusic)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	var idOfMusicInDir int
	if len(musicArray) != 0 {
		idOfMusicInDir = 0
	} else {
		fmt.Errorf("Dont have mp3 files in folder")
	}
	Player.Repeat = false
	Player.PlayPlaylist = false

	// Started music
	lenMus, _ := Player.LengthOfMusic(musicArray[idOfMusicInDir])
	lenOfMusic := float32(lenMus)

	newLayout := &SongsLayout{
		playPrevButton:        playPrevButton,
		playCurrencyButton:    playCurrencyButton,
		playNextButton:        playNextButton,
		sliderLenOfMusic:      sliderLenOfMusic,
		sliderSoundVol:        sliderSoundVol,
		optionsButton:         optionsButton,
		repeatButton:          repeatButton,
		playAllPlaylistButton: playAllPlaylistButton,
		idOfMusicInDir:        idOfMusicInDir,
		lenOfMusic:            lenOfMusic,
		Player:                Player,
		pathOfMusic:           pathOfMusic,
		MusicArray:            musicArray,
	}
	return newLayout
}

func (s *SongsLayout) Init(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		// layout.Rigid(
		// 	func(gtx layout.Context) layout.Dimensions {
		// 		return layout.Flex{
		// 			Axis: layout.Horizontal,
		// 		}.Layout(gtx,
		// 			layout.Rigid(
		// 				func(gtx layout.Context) layout.Dimensions {
		// 					optionsBtn := material.Button(th, optionsButton, "Options")
		// 					return optionsBtn.Layout(gtx)
		// 				},
		// 			),
		// 		)
		// 	},
		// ),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:    layout.Horizontal,
					Spacing: layout.Spacing(layout.Middle),
				}.Layout(gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							text := s.Player.GetName(s.MusicArray[s.idOfMusicInDir])
							return material.Body1(th, text[:len(text)-4]).Layout(gtx)
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
										material.Body1(th, fmt.Sprintf("%.0f/%.0f", s.sliderLenOfMusic.Value, s.lenOfMusic)).Layout,
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
								layout.Flexed(1, material.Slider(th, &s.sliderLenOfMusic, 0, s.lenOfMusic).Layout))
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
								layout.Flexed(1, material.Slider(th, &s.sliderSoundVol, 0, 1).Layout))
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
									playPBtn := material.Button(th, &s.repeatButton, "Repeat")
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
									playPBtn := material.Button(th, &s.playAllPlaylistButton, "Play Playlist")
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
									playPBtn := material.Button(th, &s.playPrevButton, "Prev")
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
									playPBtn := material.Button(th, &s.playCurrencyButton, "Play")
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
									playPBtn := material.Button(th, &s.playNextButton, "Next")
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
}

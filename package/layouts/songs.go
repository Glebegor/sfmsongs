package layouts

import (
	"fmt"
	"sfmsonds/package/files"
	"sfmsonds/package/music"

	"gioui.org/app"
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

func (s *SongsLayout) ListenEvents(w *app.Window) {
	if !s.sliderLenOfMusic.Dragging() && s.sliderLenOfMusic.Changed() {
		if s.Player.IsPlay {
			s.Player.StartPlayMusic(s.MusicArray[s.idOfMusicInDir], int(s.sliderLenOfMusic.Value), int(s.lenOfMusic), &s.sliderLenOfMusic, w)
		}
		if !s.Player.IsPlay {
			s.Player.IsPlay = false
		}
	}
	// Changing volume
	if s.sliderSoundVol.Dragging() {
		if s.Player.Player != nil {
			s.Player.SoundVol = float64(s.sliderSoundVol.Value)
			s.Player.SetVolume(float64(s.sliderSoundVol.Value))
		}
	}
	// RepeatButton
	if s.repeatButton.Clicked() {
		if s.Player.Repeat {
			s.Player.Repeat = false
		} else if !s.Player.Repeat {
			s.Player.Repeat = true
		}
	}
	// PlayListButton
	// if s.repeatButton.Clicked() {
	// 	if s.Player.PlayPlaylist {
	// 		s.Player.PlayPlaylist = false
	// 	} else if !s.Player.PlayPlaylist {
	// 		s.Player.PlayPlaylist = true
	// 	}
	// }

	// Play music
	if s.playCurrencyButton.Clicked() {
		if s.Player.IsPlay {
			s.Player.IsPlay = false
			s.Player.StopPlayMusic()
		} else if !s.Player.IsPlay {
			s.Player.IsPlay = true
			println("OK1")
			lenMus, _ := s.Player.LengthOfMusic(s.MusicArray[s.idOfMusicInDir])
			s.lenOfMusic = float32(lenMus)
			println(lenMus)
			println(s.MusicArray[0])
			s.Player.StartPlayMusic(s.MusicArray[s.idOfMusicInDir], int(s.sliderLenOfMusic.Value), lenMus, &s.sliderLenOfMusic, w)
		}
	}

	// Prev music
	if s.playPrevButton.Clicked() {
		if s.idOfMusicInDir > 0 {
			if s.Player.IsPlay == false {
				s.idOfMusicInDir -= 1
				lenMus, _ := s.Player.LengthOfMusic(s.MusicArray[s.idOfMusicInDir])
				s.lenOfMusic = float32(lenMus)
				s.sliderLenOfMusic.Value = 0

			} else if s.Player.IsPlay == true {
				s.idOfMusicInDir -= 1
				lenMus, err := s.Player.LengthOfMusic(s.MusicArray[s.idOfMusicInDir])
				if err != nil {
					fmt.Errorf(err.Error())
				}
				s.lenOfMusic = float32(lenMus)
				s.sliderLenOfMusic.Value = 0

				s.Player.StartPlayMusic(s.MusicArray[s.idOfMusicInDir], int(s.sliderLenOfMusic.Value), int(s.lenOfMusic), &s.sliderLenOfMusic, w)
			}
		}
	}

	// Next music
	if s.playNextButton.Clicked() {
		if s.idOfMusicInDir < len(s.MusicArray)-1 {
			if !s.Player.IsPlay {
				s.idOfMusicInDir += 1
				fmt.Println("THIS1", s.idOfMusicInDir)
				fmt.Println("THIS2", s.MusicArray)
				lenMus, _ := s.Player.LengthOfMusic(s.MusicArray[s.idOfMusicInDir])
				s.lenOfMusic = float32(lenMus)
				s.sliderLenOfMusic.Value = 0
			} else if s.Player.IsPlay {
				s.idOfMusicInDir += 1
				fmt.Println("THIS1", s.idOfMusicInDir)
				fmt.Println("THIS2", s.MusicArray)
				// fmt.Print(s.MusicArray[s.idOfMusicInDir])
				lenMus, _ := s.Player.LengthOfMusic(s.MusicArray[s.idOfMusicInDir])
				s.lenOfMusic = float32(lenMus)
				s.sliderLenOfMusic.Value = 0

				s.Player.StartPlayMusic(s.MusicArray[s.idOfMusicInDir], int(s.sliderLenOfMusic.Value), int(s.lenOfMusic), &s.sliderLenOfMusic, w)
			}
		}
	}

}
func NewSongsLayout(pathOfMusic string) *SongsLayout {
	// Buttons
	var playPrevButton widget.Clickable
	var playCurrencyButton widget.Clickable
	var playNextButton widget.Clickable
	var sliderLenOfMusic widget.Float
	var sliderSoundVol widget.Float
	var optionsButton widget.Clickable
	var repeatButton widget.Clickable
	var playAllPlaylistButton widget.Clickable

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

	Player := music.NewPlayer()
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
func (s *SongsLayout) SetSoundsArrays(pathToMusic string) {
	s.pathOfMusic = pathToMusic

	// Getting all files in dir
	MusicArray, err := files.GetMusicInFolder(s.pathOfMusic)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	s.MusicArray = MusicArray
	s.Player.Repeat = false
	s.Player.PlayPlaylist = false
}
func (s *SongsLayout) Init(gtx layout.Context, th *material.Theme) layout.Dimensions {
	s.SetSoundsArrays(s.pathOfMusic)

	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,

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
		),
		// End of buttons layout
		// Spaces
		layout.Rigid(
			layout.Spacer{Height: unit.Dp(25)}.Layout,
		), // End of spaces
	)
}

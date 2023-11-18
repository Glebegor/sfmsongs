package layouts

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type SongsListItem struct {
	Title       string
	Description string
	img         string
	MusicPath   string
}

type SongListLayout struct {
	playLists []SongsListItem
}

func (o *SongListLayout) ListenEvents(w *app.Window) {

}
func (o *SongListLayout) Init(gtx layout.Context, th *material.Theme, songsArray []string) layout.Dimensions {
	var songs []SongsListItem
	for i := 0; i < len(songsArray)-1; i++ {
		song := SongsListItem{MusicPath: songsArray[i]}
		songs = append(songs, song)
	}

	var list widget.List
	list.ScrollToEnd = true
	list.Axis = layout.Vertical

	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceAround,
				}.Layout(gtx,
					layout.Flexed(1,
						func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(16)).Layout(gtx,
								func(gtx layout.Context) layout.Dimensions {
									return list.List.Layout(gtx, len(songsArray)-1,
										func(gtx layout.Context, index int) layout.Dimensions {
											return songs[index].Layout(th, gtx)
										},
									)
								},
							)

						},
					),
				)
			},
		),
		// Spaces
		layout.Rigid(
			layout.Spacer{Height: unit.Dp(25)}.Layout,
		), // End of spaces
	)
}
func NewSongListLayout() *SongListLayout {
	return &SongListLayout{}
}

func (item SongsListItem) Layout(th *material.Theme, gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceAround,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Top:    unit.Dp(8),
					Right:  unit.Dp(8),
					Bottom: unit.Dp(8),
					Left:   unit.Dp(8),
				}.Layout(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{}.Layout(gtx,
							layout.Flexed(0.6, func(gtx layout.Context) layout.Dimensions {
								return material.Label(th, unit.Sp(10), item.MusicPath).Layout(gtx)
							}),
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									return material.Label(th, unit.Sp(10), "02:33").Layout(gtx)
								},
							),
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									var btn widget.Clickable
									return material.Button(th, &btn, "Play").Layout(gtx)
								},
							),
						)
					},
				)
			},
		),
	)
}

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
	Img         string
	MusicPath   string
	LenOfMusic  int
}

type SongListLayout struct {
	PlayLists []SongsListItem
	SongsBTNS []widget.Clickable
}

func (o *SongListLayout) ListenEvents(w *app.Window, choicer *string, songsLayer *SongsLayout) {
	for i := 0; i < len(o.SongsBTNS)-1; i++ {
		if o.SongsBTNS[i].Clicked() == true {
			*choicer = "songs"
			songsLayer.SetNewMusic(i, o.PlayLists[i].MusicPath, w)
		}
	}
}
func (o *SongListLayout) Init(gtx layout.Context, th *material.Theme, songsArray []string) layout.Dimensions {
	for i := 0; i < len(songsArray)-1; i++ {
		song := SongsListItem{MusicPath: songsArray[i]}
		o.PlayLists = append(o.PlayLists, song)
		var btn widget.Clickable
		o.SongsBTNS = append(o.SongsBTNS, btn)
	}

	var list widget.List

	list.Axis = layout.Vertical

	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Start,
		Spacing:   layout.SpaceBetween,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Vertical,
					Alignment: layout.Start,
					Spacing:   layout.SpaceAround,
				}.Layout(gtx,
					layout.Flexed(1,
						func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(16)).Layout(gtx,
								func(gtx layout.Context) layout.Dimensions {
									return list.List.Layout(gtx, len(songsArray)-1,
										func(gtx layout.Context, index int) layout.Dimensions {
											return o.PlayLists[index].Layout(th, gtx, index, o.SongsBTNS)
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

func (item SongsListItem) Layout(th *material.Theme, gtx layout.Context, id int, SongsBTN []widget.Clickable) layout.Dimensions {
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
									return layout.UniformInset(unit.Dp(10)).Layout(gtx,
										func(gtx layout.Context) layout.Dimensions {
											return material.Label(th, unit.Sp(10), "02:33").Layout(gtx)
										},
									)
								},
							),
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									return material.Button(th, &SongsBTN[id], "Play").Layout(gtx)
								},
							),
						)
					},
				)
			},
		),
	)
}

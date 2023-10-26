package layouts

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type SongsListItem struct {
	Title       string
	Description string
	img         string
	MusicArray  []string
}

type SongListLayout struct {
	playLists []SongsListItem
}

func (o *SongListLayout) ListenEvents(w *app.Window) {

}
func (o *SongListLayout) Init(gtx layout.Context, th *material.Theme) layout.Dimensions {
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
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							return material.H1(th, "SongList").Layout(gtx)
						},
					),
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							return material.H1(th, "SongList").Layout(gtx)
						},
					),
				)
			},
		),
	)
}
func NewSongListLayout() *SongListLayout {
	return &SongListLayout{}
}

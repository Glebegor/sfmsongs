package layouts

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type PlayList struct {
	Title       string
	Description string
	img         string
	MusicArray  []string
}

type PlayListsLayout struct {
	playLists []PlayList
}

func (o *PlayListsLayout) ListenEvents(w *app.Window) {

}
func (o *PlayListsLayout) Init(gtx layout.Context, th *material.Theme) layout.Dimensions {
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
							return material.H1(th, "Playlist").Layout(gtx)
						},
					),
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							return material.H1(th, "Playlist").Layout(gtx)
						},
					),
				)
			},
		),
	)
}
func NewPlayListsLayout() *PlayListsLayout {
	return &PlayListsLayout{}
}

package layouts

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type MainLayout struct {
	optionButton widget.Clickable
	//
	optionSongs     widget.Clickable
	optionPlayLists widget.Clickable
	optionThisSong  widget.Clickable

	IsOptionTrue bool
}

type Layout interface {
	Layout(gtx layout.Context, th material.Theme) layout.Dimensions
}

func NewMainLayout() *MainLayout {
	return &MainLayout{IsOptionTrue: false}
}
func (m *MainLayout) ListenEvents(w *app.Window, changedLayer string) string {
	if m.optionButton.Clicked() {
		if changedLayer == "options" {
			return "songs"
		} else {
			return "options"
		}
	}
	if m.optionSongs.Clicked() {
		return "songs"
	}
	if m.optionPlayLists.Clicked() {
		return "playList"
	}
	if m.optionThisSong.Clicked() {
		return "songList"
	}
	return changedLayer
}
func (o *MainLayout) Layout(gtx layout.Context, th *material.Theme, lay layout.Dimensions) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					// Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							optionsBtn := material.Button(th, &o.optionButton, "Options")
							return optionsBtn.Layout(gtx)
							// return ui.CreateBtnImage(gtx,&o.optionButton, "media/options.png")
						},
					),
				)
			},
		),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:    layout.Horizontal,
					Spacing: layout.SpaceAround,
				}.Layout(gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							optionsBtn := material.Button(th, &o.optionSongs, "Song")
							return optionsBtn.Layout(gtx)
						},
					),
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							optionsBtn := material.Button(th, &o.optionThisSong, "Songs")
							return optionsBtn.Layout(gtx)
						},
					),
				)
			},
		),
	)
}

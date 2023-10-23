package layouts

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type MainLayout struct {
	optionButton widget.Clickable
	IsOptionTrue bool
}

type Layout interface {
	Layout(gtx layout.Context, th material.Theme) layout.Dimensions
}

func NewMainLayout() *MainLayout {
	return &MainLayout{IsOptionTrue: false}
}
func (m *MainLayout) ListenEvents(w *app.Window) {
	if m.optionButton.Clicked() {
		fmt.Print(m.IsOptionTrue)
		m.IsOptionTrue = !m.IsOptionTrue
	}
}
func (o *MainLayout) Layout(gtx layout.Context, th *material.Theme, lay layout.Dimensions) layout.Dimensions {

	return layout.Flex{}.Layout(gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							optionsBtn := material.Button(th, &o.optionButton, "Options")
							return optionsBtn.Layout(gtx)
						},
					),
				)
			},
		),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return lay
			},
		),
	)
}

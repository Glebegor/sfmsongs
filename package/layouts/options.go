package layouts

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type OptionsLayout struct {
	LayoutMain layout.Dimensions
}

func NewOptionLayout() *OptionsLayout {
	return &OptionsLayout{}
}

//	func (o *OptionsLayout) ListenEvents(w *app.Window, songsIsTrue bool) {
//		if o.OptionBtn.Clicked() {
//			fmt.Print("EVENT")
//			// 	fmt.Print("YESSSSSSSSSSSSss")
//			// 	songsIsTrue = !songsIsTrue
//		}
//	}
func (o *OptionsLayout) Init(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:    layout.Horizontal,
				Spacing: layout.SpaceEvenly,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.H6(th, "Options").Layout(gtx)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H1(th, "Layout 2").Layout(gtx)
		}),
	)
}

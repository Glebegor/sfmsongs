package layouts

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type OptionsLayout struct {
}

type Layout interface {
	Layout(gtx layout.Context, th material.Theme) layout.Dimensions
}

func (o *OptionsLayout) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	var options widget.Clickable
	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:    layout.Horizontal,
				Spacing: layout.SpaceAround,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &options, "options").Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.H6(th, "Options").Layout(gtx)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H1(th, "Layout options2").Layout(gtx)
		}),
	)
}

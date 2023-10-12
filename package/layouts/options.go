package layouts

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type OptionsLayout struct {
	OptionBtn  widget.Clickable
	LayoutMain layout.Dimensions
}

func NewOptionLayout(gtx layout.Context, th *material.Theme) *OptionsLayout {
	return &OptionsLayout{}
}

func (o *OptionsLayout) Init(gtx layout.Context, th *material.Theme) layout.Dimensions {
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
					return material.Button(th, &o.OptionBtn, "options").Layout(gtx)
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

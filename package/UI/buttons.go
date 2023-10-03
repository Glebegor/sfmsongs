package ui

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// Function to create ui button
func CreateButton(gtx layout.Context, label string, button widget.Clickable, margins layout.Inset, th *material.Theme) layout.Dimensions {
	return margins.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			playPBtn := material.Button(th, &button, label)
			return playPBtn.Layout(gtx)
		},
	)
}
func CreateTestBut(gtx layout.Context, th *material.Theme, button widget.Clickable) layout.Dimensions {
	margins := layout.Inset{Right: unit.Dp(10), Left: unit.Dp(10)}
	return margins.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			playPBtn := material.Button(th, &button, "label")
			return playPBtn.Layout(gtx)
		},
	)
}

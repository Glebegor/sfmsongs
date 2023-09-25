package ui

import (
	"gioui.org/layout"
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

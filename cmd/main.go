package main

import (
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		// New window
		w := app.NewWindow(
			app.Title("SFMSongs"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)

		// UI variable
		var ops op.Ops

		// Play buttons
		var playPrevButton widget.Clickable
		var playCurrencyButton widget.Clickable
		var playNextButton widget.Clickable

		th := material.NewTheme()

		// listen for events in the window.
		for e := range w.Events() {

			switch e := e.(type) {

			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				// Creating layout
				layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					// Buttons play next, play current, play prev
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{
								Axis:    layout.Horizontal,
								Spacing: layout.SpaceAround,
							}.Layout(gtx,
								layout.Rigid(
									func(gtx layout.Context) layout.Dimensions {
										playPBtn := material.Button(th, &playPrevButton, "Prev")
										return playPBtn.Layout(gtx)
									},
								),
								layout.Rigid(
									func(gtx layout.Context) layout.Dimensions {
										playCBtn := material.Button(th, &playCurrencyButton, "Play")
										return playCBtn.Layout(gtx)
									},
								),
								layout.Rigid(
									func(gtx layout.Context) layout.Dimensions {
										playNBtn := material.Button(th, &playNextButton, "Next")
										return playNBtn.Layout(gtx)
									},
								),
							)
						},
					),
					layout.Rigid(
						// Bottom spacer
						layout.Spacer{Height: unit.Dp(25)}.Layout,
					),
				)
				e.Frame(gtx.Ops)
			}
		}

		// EXIT command
		os.Exit(0)
	}()
	app.Main()
}

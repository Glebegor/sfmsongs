package main

import (
	"log"
	"os"
	ui "sfmsonds/package/UI"

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
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		// EXIT command
		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window) error {

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
							Spacing: layout.Spacing(layout.Middle),
						}.Layout(gtx,
							// Prev btn
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									margins := layout.Inset{Right: unit.Dp(10), Left: unit.Dp(10)}
									return ui.CreateButton(gtx, "Prev", playPrevButton, margins, th)
								},
							),
							// Play btn
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									margins := layout.Inset{Right: unit.Dp(10), Left: unit.Dp(10)}
									return ui.CreateButton(gtx, "Play", playCurrencyButton, margins, th)
								},
							),
							// Next btn
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									margins := layout.Inset{Right: unit.Dp(10), Left: unit.Dp(10)}
									return ui.CreateButton(gtx, "Next", playNextButton, margins, th)
								},
							),
						)
					},
				), // End of buttons layout
				// Spaces
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				), // End of spaces
			)
			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}

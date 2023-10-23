package layouts

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type OptionsLayout struct {
	LayoutMain        layout.Dimensions
	musicFolderInput  widget.Editor
	musicFolderButton widget.Clickable
	MainFolder        string
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
func (o *OptionsLayout) ListenEvents(w *app.Window) {
	if o.musicFolderButton.Clicked() {
		fmt.Print(o.musicFolderInput.Text())
		o.MainFolder = o.musicFolderInput.Text()
		w.Invalidate()
	}
}
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
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Label(th, unit.Sp(16), "Path to main folder with music").Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Label(th, unit.Sp(12), o.MainFolder).Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							inputPath := material.Editor(th, &o.musicFolderInput, "Main music folder")
							return inputPath.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							inputPathBtn := material.Button(th, &o.musicFolderButton, "Main music folder")
							return inputPathBtn.Layout(gtx)
						}),
					)
				}),
			)
		}),
	)
}
func pickFolder(gtx layout.Context) (string, error) {
	var stringToPath string
	return stringToPath, nil
	// return dlg(gtx)
}

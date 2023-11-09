package layouts

import (
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
func (o *OptionsLayout) ListenEvents(songsLayer *SongsLayout, w *app.Window) {
	if o.musicFolderButton.Clicked() {
		o.MainFolder = o.musicFolderInput.Text()
		songsLayer.idOfMusicInDir = 0
		songsLayer.sliderLenOfMusic.Value = 0
		songsLayer.SetSoundsArrays(o.musicFolderInput.Text())
		w.Invalidate()
	}
	w.Invalidate()
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
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
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
			},
		),
		// Spaces
		layout.Rigid(
			layout.Spacer{Height: unit.Dp(25)}.Layout,
		), // End of spaces
	)
}

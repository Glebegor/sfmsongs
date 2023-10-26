package layouts

import (
	"fmt"
	"image"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/nfnt/resize"
)

type MainLayout struct {
	optionButton widget.Clickable
	//
	optionSongs     widget.Clickable
	optionPlayLists widget.Clickable
	optionThisSong  widget.Clickable

	IsOptionTrue bool
}

type Layout interface {
	Layout(gtx layout.Context, th material.Theme) layout.Dimensions
}

func NewMainLayout() *MainLayout {
	return &MainLayout{IsOptionTrue: false}
}
func (m *MainLayout) ListenEvents(w *app.Window, changedLayer string) string {
	if m.optionButton.Clicked() {
		if changedLayer == "options" {
			return "songs"
		} else {
			return "options"
		}
	}
	if m.optionSongs.Clicked() {
		return "songs"
	}
	if m.optionPlayLists.Clicked() {
		return "playList"
	}
	if m.optionThisSong.Clicked() {
		return "songList"
	}
	return changedLayer
}
func (o *MainLayout) Layout(gtx layout.Context, th *material.Theme, lay layout.Dimensions) layout.Dimensions {
	optionImage, err := os.Open("media/options.png")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer optionImage.Close()
	optionImageData, _, err := image.Decode(optionImage)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					// Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							optionsBtn := material.Button(th, &o.optionButton, "Options")
							paintImage(gtx, optionImageData)
							return optionsBtn.Layout(gtx)
						},
					),
				)
			},
		),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:    layout.Horizontal,
					Spacing: layout.SpaceAround,
				}.Layout(gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							optionsBtn := material.Button(th, &o.optionSongs, "Songs")
							return optionsBtn.Layout(gtx)
						},
					),
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							optionsBtn := material.Button(th, &o.optionPlayLists, "Play Lists")
							return optionsBtn.Layout(gtx)
						},
					),
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							optionsBtn := material.Button(th, &o.optionThisSong, "This Song")
							return optionsBtn.Layout(gtx)
						},
					),
				)
			},
		),
	)
}

func paintImage(gtx layout.Context, img image.Image) layout.Dimensions {
	// dims := gtx.Constraints
	// size := dims.Max

	// width, height := float32(size.X), float32(size.Y)

	// fmt.Println(x)
	// fmt.Println(y)

	// paint.PaintOp{}.Add(gtx.Ops)
	// paint.NewImageOp(img).Add(gtx.Ops)

	resizedImg := resize.Resize(uint(150), uint(150), img, resize.Lanczos3)

	imgOp := paint.NewImageOp(resizedImg)
	paint.NewImageOp(resizedImg)
	imgOp.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return layout.Dimensions{}
}

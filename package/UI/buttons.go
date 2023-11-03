package ui

import (
	"image"
	"log"
	"sfmsonds/package/files"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type BtnImage struct {
	btn widget.Clickable
}

func CreateBtnImage(gtx layout.Context, clickable *widget.Clickable, pathToImg string) layout.Dimensions {
	return material.Clickable(gtx, clickable, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(5)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			img, err := files.LoadImage(pathToImg)
			if err != nil {
				log.Fatal(err)
			}
			dim := img.Bounds().Size()
			aspectRatio := float32(dim.X) / float32(dim.Y)

			width := float32(gtx.Constraints.Max.X)
			height := width / aspectRatio

			var offset image.Point
			offset.X = int(0.5 * (float32(gtx.Constraints.Max.X) - width))
			offset.Y = int(0.5 * (float32(gtx.Constraints.Max.Y) - height))
			paint.NewImageOp(img).Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			op.Offset(offset).Add(gtx.Ops)
			return layout.Dimensions{}
		})
	})
}

func createButton(gtx layout.Context, clickable *widget.Clickable, label string) layout.Dimensions {
	return material.Clickable(gtx, clickable, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(25)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return material.Body1(nil, label).Layout(gtx)
		})
	})
}

package layouts

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type MainLayout struct {
}

type Layout interface {
	Layout(gtx layout.Context, th material.Theme) layout.Dimensions
}

func (o *MainLayout) Layout(gtx layout.Context, th *material.Theme, lay layout.Dimensions) layout.Dimensions {
	return lay
}

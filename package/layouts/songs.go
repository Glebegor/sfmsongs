package layouts

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type SongsLayout struct {
}

func NewSongsLayout(gtx layout.Context, th *material.Theme) *SongsLayout {
	newLayout := &SongsLayout{}
	return newLayout
}

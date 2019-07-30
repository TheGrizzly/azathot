package controller

import (
	"github.com/unrolled/render"
)

//Players control struct
type Player struct {
	render *render.Render
}

//NewPlayer Controller
func NewPlayer(r *render.Render) *Player {
	return &Player{
		render: r,
	}
}

package controller

import (
	"github.com/unrolled/render"
)

//Character controller struct
type Character struct {
	render *render.Render
}

//NewCharacter Controller
func NewCharacer(r *render.Render) *Character {
	return &Character{
		render: r,
	}
}

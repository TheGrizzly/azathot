package controller

import (
	"github.com/unrolled/render"
)

//users controller struct
type User struct {
	render *render.Render
}

// NewUser Controller
func NewUser(r *render.Render) *User {
	return &User{
		render: r,
	}
}

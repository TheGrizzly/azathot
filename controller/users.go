package controller

import (
	"net/http"

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

// Login handler func
func (c *User) Login(w http.ResponseWriter, req *http.Request) {
	c.render.Text(w, http.StatusOK, "Entering ...")
}

// Signup handler func
func (c *User) Signup(w http.ResponseWriter, req *http.Request) {
	c.render.Text(w, http.StatusOK, "Creating user...")
}

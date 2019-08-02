package controller

import (
	"net/http"

	"github.com/unrolled/render"
)

type Status struct {
	render *render.Render
}

//NewStatus controller
func NewStatus(r *render.Render) *Status {
	return &Status{render: r}
}

// Healthz checks external dependencies status
func (c *Status) Healthz(w http.ResponseWriter, req *http.Request) {
	//TODO: Check db Connection

	c.render.Text(w, http.StatusOK, "Friendlies")
}

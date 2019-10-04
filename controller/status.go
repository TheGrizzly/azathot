package controller

import (
	"net/http"

	"github.com/unrolled/render"
)

// Healther to check deps of Health
type Healther interface {
	IsHealthy() error
}

type HealthChecker []Healther

// Status controller struct
type Status struct {
	render   *render.Render
	healther HealthChecker
}

//NewStatus controller
func NewStatus(r *render.Render, h HealthChecker) *Status {
	return &Status{
		healther: h,
		render:   r,
	}
}

// Healthz checks external dependencies status
func (c *Status) Healthz(w http.ResponseWriter, req *http.Request) {
	if err := c.healther.IsHealthy(); err != nil {
		c.render.Text(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	c.render.Text(w, http.StatusOK, "Friendlies")
}

// IsHealthy checks deps healthiness
func (h HealthChecker) IsHealthy() error {
	for i := range h {
		if err := h[i].IsHealthy(); err != nil {
			return err
		}
	}
	return nil
}

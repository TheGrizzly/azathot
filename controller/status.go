package controller

import "net/http"

// Healthz checks external dependencies status

func (c *Player) Healthz(w http.ResponseWriter, req *http.Response) {
	//TODO: Check db Connection

	c.render.Text(w, http.StatusOK, "Friendlies")
}

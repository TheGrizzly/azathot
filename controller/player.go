package controller

import (
	cons "azathot/constant"
	"azathot/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/unrolled/render"
)

//Player usecase interface
type PlayerUsecase interface {
	GetPlayers(params *model.PlayerParams) *model.Response
}

//Players control struct
type Player struct {
	usecase PlayerUsecase
	render  *render.Render
}

//NewPlayer Controller
func NewPlayer(u PlayerUsecase, r *render.Render) *Player {
	return &Player{
		usecase: u,
		render:  r,
	}
}

//GetPlayers Handler func
func (c *Player) GetPlayers(w http.ResponseWriter, req *http.Request) {
	params, err := getPlayerParams(req)
	if err != nil {
		log.Println("error parsing getPlayer params: ", err.Error())
		c.render.Text(w, http.StatusInternalServerError, cons.UnexpectedServerError)

		return
	}

	resp := c.usecase.GetPlayers(params)
	c.render.JSON(w, resp.Code, resp.Message)
}

func getPlayerParams(req *http.Request) (*model.PlayerParams, error) {
	var params model.PlayerParams

	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return &params, nil
}

package controller

import (
	cons "azathot/constant"
	"azathot/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

//Player usecase interface
type PlayerUsecase interface {
	GetPlayers(params *model.PlayerParams) *model.Response
	GetPlayer(params *model.PlayerParams) *model.Response
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
		log.Println("error parsing getPlayer params: ", err)
		c.render.Text(w, http.StatusInternalServerError, cons.UnexpectedServerError)

		return
	}

	resp := c.usecase.GetPlayers(params)
	c.render.JSON(w, resp.Code, resp.Message)
}

func (c *Player) GetPlayer(w http.ResponseWriter, req *http.Request) {
	params, err := getPlayerParams(req)
	if err != nil {
		log.Println("error parsing getPlayer params: ", err)
		c.render.Text(w, http.StatusInternalServerError, cons.UnexpectedServerError)

		return
	}

	resp := c.usecase.GetPlayer(params)
	c.render.JSON(w, resp.Code, resp.Message)
}

func getPlayerParams(req *http.Request) (*model.PlayerParams, *model.Response) {
	var queryParams model.PlayerParams
	pathParams := mux.Vars(req)

	var playerID int
	var err error

	if pathParams["player_id"] != "" {
		playerID, err = strconv.Atoi(pathParams["player_id"])
		if err != nil {
			return nil, &model.Response{
				Code:    http.StatusBadRequest,
				Message: cons.InvalidPlayerIDMessage,
			}
		}
	}

	err = json.NewDecoder(req.Body).Decode(&queryParams)
	if err != nil {
		return nil, &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	return &model.PlayerParams{
		ID:     playerID,
		Region: queryParams.Region,
	}, nil
}

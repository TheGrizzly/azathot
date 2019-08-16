package usecase

import (
	cons "azathot/constant"
	"azathot/model"
	"log"
	"net/http"
)

//PlayerDatabase Interface
type PlayerDatabase interface {
	GetPlayers(int) ([]*model.Player, error)
}

//Player usecase
type Player struct {
	db PlayerDatabase
}

//NewPlayer usecase
func NewPlayer(db PlayerDatabase) *Player {
	return &Player{db: db}
}

func (u *Player) GetPlayers(params *model.PlayerParams) *model.Response {
	players, err := u.db.GetPlayers(params.Region)
	if err != nil {
		log.Println("error getting players: ", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	if players == nil {
		return &model.Response{
			Code:    http.StatusNotFound,
			Message: cons.PlayersNotFoundMessage,
		}
	}

	return &model.Response{
		Code: http.StatusOK,
		Message: model.PlayersResponse{
			Players: players,
		},
	}
}

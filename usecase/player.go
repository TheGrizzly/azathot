package usecase

import (
	cons "azathot/constant"
	"azathot/model"
	"log"
	"net/http"
)

//PlayerDatabase Interface
type PlayerDatabase interface {
	GetPlayers() ([]*model.Player, error)
	GetPlayerById(id int) (*model.Player, error)
	GetPlayerByName(name string) (*model.Player, error)
	InsertPlayer(p *model.Player) error
	PatchPlayer(p *model.Player) error
	DeletePlayerById(id int) error
}

//Player usecase
type Player struct {
	db PlayerDatabase
}

//NewPlayer usecase
func NewPlayer(db PlayerDatabase) *Player {
	return &Player{db: db}
}

// Get players by region func
func (u *Player) GetPlayers(params *model.PlayerParams) *model.Response {
	players, err := u.db.GetPlayers()
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

//GetPlayersByID func
func (u *Player) GetPlayer(params *model.PlayerParams) *model.Response {
	player, err := u.db.GetPlayerById(params.ID)
	if err != nil {
		log.Println("error getting player: ", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	if player == nil {
		return &model.Response{
			Code:    http.StatusNotFound,
			Message: cons.PlayerNotFoundMessage,
		}
	}

	return &model.Response{
		Code: http.StatusOK,
		Message: model.PlayersResponse{
			Player: player,
		},
	}
}

//PostPlayer func
func (u *Player) PostPlayer(params *model.PlayerParams) *model.Response {
	player, err := u.db.GetPlayerByName(params.ReqPlayer.Name)
	if err != nil {
		log.Println("error getting product by name:", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	if player != nil {
		return &model.Response{
			Code:    http.StatusBadRequest,
			Message: cons.PlayerAlreadyExistsMessage,
		}
	}

	err = u.db.InsertPlayer(params.ReqPlayer)
	if err != nil {
		log.Println("error inserting user:", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	player, err = u.db.GetPlayerByName(params.ReqPlayer.Name)
	if err != nil {
		log.Println("error getting product by name:", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	if player == nil {
		return &model.Response{
			Code:    http.StatusBadRequest,
			Message: cons.PlayerNotFoundMessage,
		}
	}

	return &model.Response{
		Code: http.StatusOK,
		Message: model.PlayersResponse{
			Player: player,
		},
	}
}

//PatchPlayer func
func (u *Player) PatchPlayer(params *model.PlayerParams) *model.Response {
	player, err := u.db.GetPlayerById(params.ID)
	if err != nil {
		log.Println("error getting player by id: ", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	if player == nil {
		return &model.Response{
			Code:    http.StatusBadRequest,
			Message: cons.PlayerNotFoundMessage,
		}
	}

	err = u.db.PatchPlayer(params.ReqPlayer)
	if err != nil {
		log.Println("error patching player: ", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	player, err = u.db.GetPlayerById(params.ID)
	if err != nil {
		log.Println("error getting player by id: ", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	if player == nil {
		return &model.Response{
			Code:    http.StatusBadRequest,
			Message: cons.PlayerNotFoundMessage,
		}
	}

	return &model.Response{
		Code: http.StatusOK,
		Message: model.PlayersResponse{
			Player: player,
		},
	}
}

func (u *Player) DeletePlayerById(params *model.PlayerParams) *model.Response {
	player, err := u.db.GetPlayerById(params.ID)
	if err != nil {
		log.Println("error getting player by id: ", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	if player == nil {
		return &model.Response{
			Code:    http.StatusBadRequest,
			Message: cons.PlayerNotFoundMessage,
		}
	}

	err = u.db.DeletePlayerById(params.ID)
	if err != nil {
		log.Println("error deleting player: ", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	return &model.Response{
		Code:    http.StatusOK,
		Message: cons.PlayerDeletedMessage,
	}
}

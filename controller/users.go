package controller

import (
	cons "azathot/constant"
	"azathot/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/unrolled/render"
)

// UserUsecase interface
type UserUsecase interface {
	Signup(params *model.UserParams) *model.Response
}

//users controller struct
type User struct {
	usecase UserUsecase
	render  *render.Render
}

// NewUser Controller
func NewUser(u UserUsecase, r *render.Render) *User {
	return &User{
		usecase: u,
		render:  r,
	}
}

// Login handler func
func (c *User) Login(w http.ResponseWriter, req *http.Request) {
	c.render.Text(w, http.StatusOK, "Entering ...")
}

// Signup handler func
func (c *User) Signup(w http.ResponseWriter, req *http.Request) {
	params, err := getUserParams(req)
	if err != nil {
		log.Println("error parsing signup params:", err.Error())
		c.render.Text(w, http.StatusInternalServerError, cons.UnexpectedServerError)

		return
	}

	resp := c.usecase.Signup(params)
	c.render.JSON(w, resp.Code, resp.Message)
}

func getUserParams(req *http.Request) (*model.UserParams, error) {
	var params model.UserParams

	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return &params, nil
}

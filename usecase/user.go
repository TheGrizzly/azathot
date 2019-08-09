package usecase

import (
	cons "azathot/constant"
	"azathot/model"
	"log"
	"net/http"
)

// UserDatabase interace
type UserDatabase interface {
	GetUserByEmail(email string) (*model.User, error)
	InsertUser(email, password string) error
}

// Crypter interface
type Crypter interface {
	EncryptPassword(passwd string) string
	DoesPasswordsMatch(reqPasswd, dbPasswd string) bool
}

type JWT interface {
	Generate(email string, isAdmin bool) (string, error)
}

// User usecas struct
type User struct {
	db      UserDatabase
	crypter Crypter
	jwt     JWT
}

// NewUser usecase
func NewUser(db UserDatabase, c Crypter, j JWT) *User {
	return &User{
		db:      db,
		crypter: c,
		jwt:     j,
	}
}

//Signup users flow
func (u *User) Signup(params *model.UserParams) *model.Response {
	user, err := u.db.GetUserByEmail(params.Email)
	if err != nil {
		log.Println("error getting user by email: ", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	if user != nil {
		return &model.Response{
			Code:    http.StatusBadRequest,
			Message: cons.UserAlreadyExistsError,
		}
	}

	err = u.db.InsertUser(params.Email, u.crypter.EncryptPassword(params.Password))
	if err != nil {
		log.Println("error inserting users: ", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	return &model.Response{
		Code:    http.StatusOK,
		Message: cons.UserRegisteredMessage,
	}
}

//Login flow for users
func (u *User) Login(params *model.UserParams) *model.Response {
	user, err := u.db.GetUserByEmail(params.Email)
	if err != nil {
		log.Println("error getting user by email:", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	if user == nil || !u.crypter.DoesPasswordsMatch(params.Password, user.Password) {
		return &model.Response{
			Code:    http.StatusUnauthorized,
			Message: cons.InvalidCredentialMessage,
		}
	}

	token, err := u.jwt.Generate(user.Email, user.IsAdmin)
	if err != nil {
		log.Println("error generating token:", err.Error())

		return &model.Response{
			Code:    http.StatusInternalServerError,
			Message: cons.UnexpectedServerError,
		}
	}

	return &model.Response{
		Code: http.StatusOK,
		Message: model.LoginResponse{
			Message: cons.LoginSucessfulMessage,
			Token:   token,
		},
	}
}

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
}

// User usecas struct
type User struct {
	db      UserDatabase
	crypter Crypter
}

// NewUser usecase
func NewUser(db UserDatabase, c Crypter) *User {
	return &User{
		db:      db,
		crypter: c,
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

package crypt

import (
	"azathot/config"

	"golang.org/x/crypto/bcrypt"
)

// Service for crypt stuff
type Service struct {
	cost int
}

// New crypt service
func New(ac *config.App) *Service {
	return &Service{cost: ac.CryptCost}
}

// EncryptPassword for saving in the DB
func (s *Service) EncryptPassword(passwd string) string {
	bytePasswd, _ := bcrypt.GenerateFromPassword([]byte(passwd), s.cost)
	return string(bytePasswd)
}

// DoesPasswordsMatch for request password and db password
func (s *Service) DoesPasswordsMatch(reqPasswd, dbPasswd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPasswd), []byte(reqPasswd))
	return err == nil
}

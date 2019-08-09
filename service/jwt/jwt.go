package jwt

import (
	"azathot/config"
	"fmt"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// Service for JWT service
type Service struct {
	key             []byte
	tokenExpiration time.Duration
}

// New JWT Service
func New(ac *config.App) *Service {
	duration, _ := time.ParseDuration(fmt.Sprintf("%dm", ac.JWTExpiration))

	return &Service{
		key:             []byte(ac.JWTKey),
		tokenExpiration: duration,
	}
}

// Generate JWT
func (s *Service) Generate(email string, isAdmin bool) (string, error) {
	token := jwtgo.New(jwtgo.SigningMethodHS256)

	claims := token.Claims.(jwtgo.MapClaims)
	claims["email"] = email
	claims["isAdmin"] = isAdmin
	claims["exp"] = time.Now().Add(s.tokenExpiration).Unix()

	return token.SignedString(s.key)
}

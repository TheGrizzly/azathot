package middleware

import (
	cons "azathot/constant"
	"net/http"

	"github.com/unrolled/render"
)

// JWT inteface
type JWT interface {
	Validate(token string) bool
}

// Middleware struct
type Middleware struct {
	jwt    JWT
	render *render.Render
}

func New(j JWT, r *render.Render) *Middleware {
	return &Middleware{
		jwt:    j,
		render: r,
	}
}

func (m *Middleware) ValidateJWT(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if !m.jwt.Validate(req.Header.Get("Token")) {
			m.render.Text(w, http.StatusUnauthorized, cons.InvalidTokenMessage)
			return
		}

		handler(w, req)
	}
}

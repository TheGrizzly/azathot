package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"

	"azathot/router/mocks"
)

func TestSetup(t *testing.T) {
	testCases := []struct {
		desc                string
		handlerName         string
		handlerMethod       string
		status              int
		expectStatusCall    bool
		expectPlayerCall    bool
		expectUserCall      bool
		expectJWTValidation bool
	}{
		{
			desc:             "/healthz",
			handlerName:      "Healthz",
			handlerMethod:    "Get",
			expectStatusCall: true,
			status:           http.StatusOK,
		},
		{
			desc:           "/login",
			handlerName:    "Login",
			handlerMethod:  "Post",
			expectUserCall: true,
			status:         http.StatusOK,
		},
		{
			desc:           "/signup",
			handlerName:    "Signup",
			handlerMethod:  "Post",
			expectUserCall: true,
			status:         http.StatusOK,
		},
		{
			desc:             "/players",
			handlerName:      "GetPlayers",
			handlerMethod:    "GetAll",
			expectPlayerCall: true,
			status:           http.StatusOK,
		},
		{
			desc:             "/players/a_id1",
			handlerName:      "GetPlayer",
			handlerMethod:    "GetSingle",
			expectPlayerCall: true,
			status:           http.StatusOK,
		},
		{
			desc:                "/admin/players",
			handlerName:         "PostPlayer",
			handlerMethod:       "Post",
			expectPlayerCall:    true,
			expectJWTValidation: true,
			status:              http.StatusOK,
		},
		{
			desc:                "/admin/players/a_id1",
			handlerName:         "PatchPlayer",
			handlerMethod:       "Patch",
			expectPlayerCall:    true,
			expectJWTValidation: true,
			status:              http.StatusOK,
		},
		{
			desc:                "/admin/players/a_id1",
			handlerName:         "DeletePlayerById",
			handlerMethod:       "Delete",
			expectPlayerCall:    true,
			expectJWTValidation: true,
			status:              http.StatusOK,
		},
		{
			desc:   "/badRequest",
			status: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			statusController := &mocks.StatusController{}
			userController := &mocks.UserController{}
			playerController := &mocks.PlayerController{}
			jwtMiddleware := &mocks.JWTMiddelware{}

			switch {
			case tc.expectStatusCall:
				statusController.On(tc.handlerName, mock.Anything, mock.Anything).
					Run(func(args mock.Arguments) {
						args[0].(http.ResponseWriter).WriteHeader(http.StatusOK)
					})
			case tc.expectUserCall:
				userController.On(tc.handlerName, mock.Anything, mock.Anything).
					Run(func(args mock.Arguments) {
						args[0].(http.ResponseWriter).WriteHeader(http.StatusOK)
					})
			case tc.expectPlayerCall:
				playerController.On(tc.handlerName, mock.Anything, mock.Anything).
					Run(func(args mock.Arguments) {
						args[0].(http.ResponseWriter).WriteHeader(http.StatusOK)
					})
			}

			req, err := http.NewRequest("GET", tc.desc, nil)

			switch tc.handlerMethod {
			case "GetAll":
				req, err = http.NewRequest("GET", tc.desc, nil)
			case "Post":
				req, err = http.NewRequest("POST", tc.desc, nil)
			case "Patch":
				req, err = http.NewRequest("PATCH", tc.desc, nil)
			case "Delete":
				req, err = http.NewRequest("DELETE", tc.desc, nil)
			}

			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			var arg http.HandlerFunc
			arg = nil

			switch tc.handlerName {
			case "PostPlayer":
				arg = playerController.PostPlayer
			case "PatchPlayer":
				arg = playerController.PatchPlayer
			case "DeletePlayerById":
				arg = playerController.DeletePlayerById
			}

			if tc.expectJWTValidation {
				jwtMiddleware.On("ValidateJWT", mock.AnythingOfType("http.HandlerFunc")).Return(arg)
			} else {
				jwtMiddleware.On("ValidateJWT", mock.Anything).Return(nil)
			}

			r := GetRouter(statusController, userController, playerController, jwtMiddleware)

			r.ServeHTTP(rr, req)

			statusController.AssertExpectations(t)
			userController.AssertExpectations(t)
			playerController.AssertExpectations(t)
			jwtMiddleware.AssertExpectations(t)

			assert.Equal(t, tc.status, rr.Code)
		})
	}
}

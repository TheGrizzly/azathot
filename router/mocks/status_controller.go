// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import http "net/http"
import mock "github.com/stretchr/testify/mock"

// StatusController is an autogenerated mock type for the StatusController type
type StatusController struct {
	mock.Mock
}

// Healthz provides a mock function with given fields: w, req
func (_m *StatusController) Healthz(w http.ResponseWriter, req *http.Request) {
	_m.Called(w, req)
}
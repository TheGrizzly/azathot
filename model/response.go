package model

//response type struct
type Response struct {
	Code    int
	Message interface{}
}

// Login response-
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

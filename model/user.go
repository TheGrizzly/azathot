package model

// user params for related flows like login or signup
type UserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User struct for app usage
type User struct {
	ID       int64
	Email    string
	Password string
	IsAdmin  bool
	IdPlayer int64
}

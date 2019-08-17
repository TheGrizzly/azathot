package model

// PlayerParams
type PlayerParams struct {
	Region int `json:"region"`
	ID     int `json:"id"`
}

// Player model
type Player struct {
	ID          int64
	Name        string
	Tag         string
	IdMain      int64
	SmashggUser string
	NumColor    int64
	IdRegion    int64
}

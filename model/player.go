package model

// PlayerParams
type PlayerParams struct {
	ID        int     `json:"id"`
	ReqPlayer *Player `json:"player"`
}

// Player model
type Player struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	IdMain      int64  `json:"id_main"`
	SmashggUser string `json:"smashgg_user"`
	NumColor    int64  `json:"num_color"`
	IdRegion    int64  `json:"id_region"`
}

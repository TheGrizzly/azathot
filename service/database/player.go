package database

import (
	"azathot/model"
	"database/sql"
)

const (
	getPlayersQuery = `
		SELECT 
			id,
			name,
			tag,
			id_main,
			smashgg_user,
			num_color,
			id_region
		FROM
			players
		WHERE
			id_region = ?
	`
)

type Player struct {
	ID          sql.NullInt64  `db:"id"`
	Name        sql.NullString `db:"name"`
	Tag         sql.NullString `db:"tag"`
	IdMain      sql.NullInt64  `db:"id_main"`
	SmashggUser sql.NullString `db:"smashgg_user"`
	NumColor    sql.NullInt64  `db:"num_color"`
	IdRegion    sql.NullInt64  `db:"id_region"`
}

type Players []Player

func (s *Service) GetPlayers(IdRegion int) ([]*model.Player, error) {
	var dbPlayers Players

	rows, err := s.db.Query(getPlayersQuery, IdRegion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tempPlayer := Player{}

		err := rows.Scan(
			&tempPlayer.ID,
			&tempPlayer.Name,
			&tempPlayer.Tag,
			&tempPlayer.IdMain,
			&tempPlayer.SmashggUser,
			&tempPlayer.NumColor,
			&tempPlayer.IdRegion,
		)
		if err != nil {
			return nil, err
		}

		dbPlayers = append(dbPlayers, tempPlayer)
	}

	if len(dbPlayers) == 0 {
		return nil, nil
	}

	return dbPlayers.ToModel(), nil
}

func (p Players) ToModel() []*model.Player {
	mPlayers := []*model.Player{}

	for n := range p {
		mPlayers = append(mPlayers, &model.Player{
			ID:          p[n].ID.Int64,
			Name:        p[n].Name.String,
			Tag:         p[n].Tag.String,
			IdMain:      p[n].IdMain.Int64,
			SmashggUser: p[n].SmashggUser.String,
			NumColor:    p[n].NumColor.Int64,
			IdRegion:    p[n].IdRegion.Int64,
		})
	}

	return mPlayers
}

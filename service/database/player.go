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
	`

	getPlayersByRegionQuery = getPlayersQuery + `
		WHERE id_region = ?
	`

	getPlayerByIdQuery = getPlayersQuery + `
		WHERE id = ?
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

	rows, err := s.db.Query(getPlayersByRegionQuery, IdRegion)
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

func (s *Service) GetPlayerById(id int) (*model.Player, error) {
	var dbPlayer Player

	err := s.db.QueryRow(getPlayerByIdQuery, id).Scan(
		&dbPlayer.ID,
		&dbPlayer.Name,
		&dbPlayer.Tag,
		&dbPlayer.IdMain,
		&dbPlayer.SmashggUser,
		&dbPlayer.NumColor,
		&dbPlayer.IdRegion,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return dbPlayer.toModel(), err
}

func (p Player) toModel() *model.Player {
	return &model.Player{
		ID:          p.ID.Int64,
		Name:        p.Name.String,
		Tag:         p.Tag.String,
		IdMain:      p.IdMain.Int64,
		SmashggUser: p.SmashggUser.String,
		NumColor:    p.NumColor.Int64,
		IdRegion:    p.IdRegion.Int64,
	}
}

func (p Players) ToModel() []*model.Player {
	mPlayers := []*model.Player{}

	for n := range p {
		mPlayers = append(mPlayers, p[n].toModel())
	}

	return mPlayers
}

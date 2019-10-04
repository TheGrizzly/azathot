package database

import (
	"azathot/model"
	"database/sql"
)

const (
	getUserByEmailQuery = `
		SELECT
			u.id,
			u.email,
			u.password,
			u.is_admin,
			(SELECT p.id_player FROM player_pivot AS p WHERE p.id_user = u.id) as id_player
		FROM users as u
		WHERE u.email = ?
	`
	insertUserQuery = `
		INSERT INTO
			users (email, password, is_admin)
		VALUES
			(?, ?, false)
	`
)

type User struct {
	ID       sql.NullInt64  `db:"id"`
	Email    sql.NullString `db:"email"`
	Password sql.NullString `db:"password"`
	IsAdmin  sql.NullBool   `db:"is_admin"`
	IdPlayer sql.NullInt64  `db:"id_player"`
}

// GetUserByEmail - if it was not found, it returns nil and error
func (s *Service) GetUserByEmail(email string) (*model.User, error) {
	var dbUser User

	err := s.db.QueryRow(getUserByEmailQuery, email).Scan(
		&dbUser.ID,
		&dbUser.Email,
		&dbUser.Password,
		&dbUser.IsAdmin,
		&dbUser.IdPlayer,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return dbUser.toModel(), nil
}

// InsertUser registers a new user into the database
func (s *Service) InsertUser(email, password string) error {
	stmt, err := s.db.Prepare(insertUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, password)
	if err != nil {
		return err
	}

	return nil
}

func (u User) toModel() *model.User {
	return &model.User{
		ID:       u.ID.Int64,
		Email:    u.Email.String,
		Password: u.Password.String,
		IsAdmin:  u.IsAdmin.Bool,
		IdPlayer: u.IdPlayer.Int64,
	}
}

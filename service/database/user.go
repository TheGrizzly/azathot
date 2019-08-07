package database

import (
	"azathot/model"
	"database/sql"
)

const (
	getUserByEmailQuery = `
		SELECT
			id,
			email,
			password,
			is_admin,
			id_player
		FROM users
		WHERE email = ?
	`
	insertUserQuery = `
		INSERT INTO
			users (email, password)
		VALUES
			(?, ?)
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

	err := s.db.QueryRow(getUserByEmailQuery, email).Scan(&dbUser.ID, &dbUser.Email, &dbUser.Password, &dbUser.IsAdmin, &dbUser.IdPlayer)
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

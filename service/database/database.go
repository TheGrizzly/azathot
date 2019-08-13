package database

import (
	"azathot/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
)

// Service for database calls
type Service struct {
	db *sqlx.DB
}

// New database service struct
func New(ac *config.App) (*Service, error) {
	db, err := sqlx.Open(ac.DBDriver, fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		ac.DBUser,
		ac.DBPassword,
		ac.DBHost,
		ac.DBPort,
		ac.DBName,
	))
	if err != nil {
		return nil, err
	}

	return &Service{db: db}, nil
}

//isHealthy check DB connection
func (s *Service) IsHealthy() error {
	return s.db.Ping()
}

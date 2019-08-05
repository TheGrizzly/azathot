package database

import (
	"azathot/config"
	"database/sql"
	"fmt"
)

// Service for database calls
type Service struct {
	db *sql.DB
}

// New database service struct
func New(ac *config.App) (*Service, error) {
	db, err := sql.Open(ac.DBDriver, fmt.Sprintf(
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

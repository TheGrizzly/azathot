package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
)

//App struct to be used globally
type App struct {
	DBDriver   string `json:"db_driver"`
	DBName     string `json:"db_name"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	CryptCost  int    `json:"crypt_cost"`
}

const (
	configPathFile = "./config.json"
)

//Load configuration from file
func LoadFromConfigFile() (*App, error) {
	var app App

	configFile, err := ioutil.ReadFile(filepath.Clean(configPathFile))
	if err != nil {
		return nil, errors.New("failed to read file")
	}

	err = json.Unmarshal(configFile, &app)
	if err != nil {
		return nil, errors.New("failed to unmarshal data from json")
	}

	return &app, nil
}

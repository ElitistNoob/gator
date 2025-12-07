package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilepath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home location: %w", err)
	}

	path := filepath.Join(home, "/workspace/github/gator/", configFileName)
	return path, nil
}

func Read() (*Config, error) {
	path, err := getConfigFilepath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file at location %v: %w", path, err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding content from file %v: %w", file, err)
	}

	return &config, nil
}

func Write(cfg Config) error {
	path, err := getConfigFilepath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(&cfg); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.Current_user_name = name

	if err := Write(*cfg); err != nil {
		return err
	}

	return nil
}

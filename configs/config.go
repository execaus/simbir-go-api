package configs

import (
	"encoding/json"
	"github.com/execaus/exloggo"

	"os"
)

const configFilePath = "configs/config.json"

type Config struct {
	Postgres *PostgresConfig `json:"postgres,omitempty"`
	Server   *ServerConfig   `json:"server,omitempty"`
}

type PostgresConfig struct {
	Host    string `json:"host,omitempty"`
	Port    string `json:"port,omitempty"`
	DBName  string `json:"DBName,omitempty"`
	SSLMode string `json:"SSLMode,omitempty"`
}

type ServerConfig struct {
	Port string `json:"port,omitempty"`
}

func LoadConfig() *Config {
	file, _ := os.Open(configFilePath)
	defer closeFile(file)

	decoder := json.NewDecoder(file)

	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		exloggo.Fatalf("parse config file: %s", err.Error())
	}

	return &config
}

func closeFile(file *os.File) {
	_ = file.Close()
}

package models

import (
	"github.com/caarlos0/env/v9"
	"github.com/execaus/exloggo"
	"github.com/joho/godotenv"
)

type Environment struct {
	PostgresUser     string `env:"POSTGRES_USER,required,notEmpty"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required,notEmpty"`
}

const envFilePath = ".env.local"

func LoadEnv() *Environment {
	environment := Environment{}
	if err := godotenv.Load(envFilePath); err != nil {
		exloggo.Warningf("load file not found (%s), environment variables load from environment", err.Error())
	}
	if err := env.Parse(&environment); err != nil {
		exloggo.Fatalf(`environment variables load from environment: %s`, err.Error())
	}

	return &environment
}

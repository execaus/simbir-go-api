package repository

import (
	"database/sql"
	"fmt"
	"github.com/execaus/exloggo"
	_ "github.com/lib/pq" //nolint:nolintlint,revive
	"simbir-go-api/configs"
	"simbir-go-api/models"
	"simbir-go-api/queries"
)

const dbDriverName = "postgres"

func NewBusinessDatabase(env *models.Environment, config *configs.Config) *queries.Queries {
	conn := getConnectDatabase(env, config)
	db := queries.New(conn)
	return db
}

func getConnectDatabase(env *models.Environment, config *configs.Config) *sql.DB {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Postgres.Host,
		config.Postgres.Port,
		env.PostgresUser,
		env.PostgresPassword,
		config.Postgres.DBName,
		config.Postgres.SSLMode)
	db, err := sql.Open(dbDriverName, connString)
	if err != nil {
		exloggo.Fatalf(`database open connect: %s`, err.Error())
	}
	if err = db.Ping(); err != nil {
		exloggo.Fatalf(`database open connect: %s`, err.Error())
	}
	return db
}

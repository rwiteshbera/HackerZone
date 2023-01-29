package database

import (
	_ "database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rwiteshbera/HackZone/api"
)

func Connect(server *api.Server) (*sqlx.DB, error) {
	db, err := sqlx.Connect(server.Config.DB_DRIVER, server.Config.DB_CONNECTION_STRING)
	if err != nil {
		return nil, err
	}
	return db, nil
}

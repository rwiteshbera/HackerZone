package database

import (
	"database/sql"
	_ "database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rwiteshbera/HackerZone/api"
	"time"
)

func Connect(server *api.Server) (*sql.DB, error) {
	db, err := sql.Open(server.Config.DB_DRIVER, server.Config.DB_CONNECTION_STRING)
	if err != nil {
		return nil, errors.Wrap(err, "package : database")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	return db, nil
}

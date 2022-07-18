package storage

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

const (
	usersTable = "content.users"
	ordersTable = "content.orders"
	withdrawalTable = "content.withdrawals"
)

func NewPostgresDB(uri string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		return nil, nil
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	migration, err := migrate.New("file://migrations", uri)
	if err != nil {
		return nil, err
	}
	if err := migration.Up(); err != nil {
		log.Info().Msgf("migrations: %s", err.Error())
	}

	return db, nil
}

package postgres

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Username string
	Port     string
	Password string
	DBName   string
	SSL      string
}

func NewPostgresDB(ctx context.Context, connString string, sourceURL string) (*pgx.Conn, error) {
	migration(connString, sourceURL)
	db, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migration(connString string, sourceURL string) {
	m, err := migrate.New(
		"file:/"+sourceURL,
		connString,
	)
	if err != nil {
		logrus.Fatalf("error reading migration: %s", err.Error())
	}
	if err := m.Up(); err != nil {
		logrus.Fatalf("error up migration: %s", err.Error())
	}
}

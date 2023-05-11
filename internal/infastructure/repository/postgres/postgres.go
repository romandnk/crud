package postgres

import (
	"context"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

const (
	tasksTable = "tasks"
)

type Config struct {
	Host     string
	Username string
	Port     string
	Password string
	DBName   string
	SSL      string
}

func NewPostgresDB(ctx context.Context, connString string) (*pgx.Conn, error) {
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

package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Config struct {
	Host     string
	Username string
	Port     string
	Password string
	DBName   string
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

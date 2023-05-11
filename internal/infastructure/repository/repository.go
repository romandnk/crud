package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/romandnk/crud/internal/entities"
	"github.com/romandnk/crud/internal/infastructure/repository/postgres"
)

type Tasker interface {
	Create(ctx context.Context, task entities.Task) (int, error)
	GetAll(ctx context.Context) ([]entities.Task, error)
	GetById(ctx context.Context, id int) (entities.Task, error)
	Update(ctx context.Context, id int, task entities.Task) (entities.Task, error)
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	Tasker
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Tasker: postgres.NewTaskPostgres(db),
	}
}

package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/romandnk/crud/internal/entities/task"
	"github.com/romandnk/crud/internal/infastructure/repository/postgres"
)

type Task interface {
	Create(ctx context.Context, task task.Task) (int, error)
	GetAll(ctx context.Context) ([]task.Task, error)
	GetById(ctx context.Context, id int) (task.Task, error)
	Update(ctx context.Context, id int, task task.Task) (task.Task, error)
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	Task
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Task: postgres.NewTaskPostgres(db),
	}
}

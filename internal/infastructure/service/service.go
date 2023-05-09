package service

import (
	"context"
	"github.com/romandnk/crud/internal/entities/task"
	"github.com/romandnk/crud/internal/infastructure/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock_db.go
type Task interface {
	Create(ctx context.Context, task task.Task) (int, error)
	GetAll(ctx context.Context) ([]task.Task, error)
	GetById(ctx context.Context, id int) (task.Task, error)
	Update(ctx context.Context, id int, task task.Task) (task.Task, error)
	Delete(ctx context.Context, id int) error
}

type Service struct {
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Task: NewTaskService(repos.Task),
	}
}

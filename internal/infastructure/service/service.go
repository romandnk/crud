package service

import (
	"context"
	"github.com/romandnk/crud/internal/entities"
	"github.com/romandnk/crud/internal/infastructure/repository"
)

//go:generate mockgen -source=service.go -destination=../mocks/mock_db.go
type Tasker interface {
	Create(ctx context.Context, task entities.Task) (int, error)
	GetAll(ctx context.Context) ([]entities.Task, error)
	GetById(ctx context.Context, id int) (entities.Task, error)
	Update(ctx context.Context, id int, task entities.Task) (entities.Task, error)
	Delete(ctx context.Context, id int) error
}

type Service struct {
	Tasker
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Tasker: NewTaskService(repos.Tasker),
	}
}

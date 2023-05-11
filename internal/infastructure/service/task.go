package service

import (
	"context"
	"github.com/romandnk/crud/internal/entities"
	"github.com/romandnk/crud/internal/infastructure/repository"
)

type TaskService struct {
	repo repository.Tasker
}

func NewTaskService(repo repository.Tasker) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (t *TaskService) Create(ctx context.Context, task entities.Task) (int, error) {
	return t.repo.Create(ctx, task)
}

func (t *TaskService) GetAll(ctx context.Context) ([]entities.Task, error) {
	return t.repo.GetAll(ctx)
}
func (t *TaskService) GetById(ctx context.Context, id int) (entities.Task, error) {
	return t.repo.GetById(ctx, id)
}
func (t *TaskService) Update(ctx context.Context, id int, task entities.Task) (entities.Task, error) {
	return t.repo.Update(ctx, id, task)
}
func (t *TaskService) Delete(ctx context.Context, id int) error {
	return t.repo.Delete(ctx, id)
}

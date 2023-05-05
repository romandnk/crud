package service

import (
	"context"
	"github.com/romandnk/crud/internal/entities/task"
	"github.com/romandnk/crud/internal/infastructure/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (t *TaskService) Create(ctx context.Context, task task.Task) (int, error) {
	return t.repo.Create(ctx, task)
}

func (t *TaskService) GetAll(ctx context.Context) ([]task.Task, error) {
	return t.repo.GetAll(ctx)
}
func (t *TaskService) GetById(ctx context.Context, id int) (task.Task, error) {
	return t.repo.GetById(ctx, id)
}
func (t *TaskService) Update(ctx context.Context, id int, task task.Task) (task.Task, error) {
	return t.repo.Update(ctx, id, task)
}
func (t *TaskService) Delete(ctx context.Context, id int) error {
	return t.repo.Delete(ctx, id)
}

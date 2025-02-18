package service

import (
	"context"
	"todo/internal/controllers/rest/handlers/domain"
)

type TaskService interface {
	AddTask(ctx context.Context, title, description, status string) (int, error)
	GetTask(ctx context.Context, id int) (*domain.Task, error)
	GetTasks(ctx context.Context, limit, page int, order string) ([]*domain.Task, error)
	RemoveTask(ctx context.Context, id int) error
	UpdateTask(ctx context.Context, id int, title, description, status string) error
}

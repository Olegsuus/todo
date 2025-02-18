package storage

import (
	"context"
	"todo/internal/service/models"
	"todo/internal/storage/entity"
)

type TaskStorage interface {
	AddTask(ctx context.Context, taskEntity entity.TaskEntity) (*models.Task, error)
	GetTask(ctx context.Context, id int) (*models.Task, error)
	GetTasks(ctx context.Context, limit, offset int, order string) ([]*models.Task, error)
	RemoveTask(ctx context.Context, id int) error
	UpdateTask(ctx context.Context, id int, taskEntity entity.TaskEntity) error
}

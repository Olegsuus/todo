package service

import (
	"context"
	"todo/internal/storage/entity"
)

func (s *TaskService) AddTask(ctx context.Context, title, description, status string) (int, error) {
	entityTask := entity.TaskEntity{
		Title:       title,
		Description: description,
		Status:      status,
	}

	task, err := s.taskStorage.AddTask(ctx, entityTask)
	if err != nil {
		return -1, err
	}

	return task.ID, nil
}

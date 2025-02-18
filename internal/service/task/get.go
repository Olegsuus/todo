package service

import (
	"context"
	"todo/internal/controllers/rest/handlers/domain"
)

func (s *TaskService) GetTask(ctx context.Context, id int) (*domain.Task, error) {
	task, err := s.taskStorage.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	return modelsToDomain(task), nil
}

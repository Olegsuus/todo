package service

import (
	"context"
	"todo/internal/controllers/rest/handlers/domain"
	"todo/pkg/utils"
)

func (s *TaskService) GetTasks(ctx context.Context, limit, page int, order string) ([]*domain.Task, error) {
	if order != "asc" && order != "desc" {
		order = "desc"
	}
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 1
	}
	offset := (page - 1) * limit

	modelsTasks, err := s.taskStorage.GetTasks(ctx, limit, offset, order)
	if err != nil {
		return nil, err
	}

	tasks := utils.MapAsync(modelsTasks, modelsToDomain)

	return tasks, nil
}

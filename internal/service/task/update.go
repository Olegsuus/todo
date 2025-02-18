package service

import (
	"context"
	"time"
)

func (s *TaskService) UpdateTask(ctx context.Context, id int, title, description, status string) error {
	modelsTask, err := s.taskStorage.GetTask(ctx, id)
	if err != nil {
		return err
	}

	if title != "" {
		modelsTask.Title = title
	}

	if description != "" {
		modelsTask.Description = description
	}

	if status != "" {
		modelsTask.Status = status
	}

	modelsTask.UpdatedAt = time.Now()

	taskEntity := modelsToEntity(modelsTask)

	if err := s.taskStorage.UpdateTask(ctx, id, *taskEntity); err != nil {
		return err
	}

	return nil
}

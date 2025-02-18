package service

import (
	"context"
)

func (s *TaskService) RemoveTask(ctx context.Context, id int) error {
	if err := s.taskStorage.RemoveTask(ctx, id); err != nil {
		return err
	}

	return nil
}

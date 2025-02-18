package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"todo/internal/service/models"
	"todo/internal/storage/entity"
	apperror "todo/pkg/errors"
)

func (s *TaskStorage) AddTask(ctx context.Context, taskEntity entity.TaskEntity) (*models.Task, error) {
	s.l.Info("создание новой задачи:", taskEntity)

	query, args, err := squirrel.
		Insert("tasks").
		Columns("title", "description", "status").
		Values(taskEntity.Title, taskEntity.Description, taskEntity.Status).
		Suffix("RETURNING id, title, description, status, created_at, updated_at").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, apperror.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при создании задачи",
		}
	}

	err = s.db.QueryRow(ctx, query, args...).Scan(
		&taskEntity.ID,
		&taskEntity.Title,
		&taskEntity.Description,
		&taskEntity.Status,
		&taskEntity.CreatedAt,
		&taskEntity.UpdatedAt,
	)
	if err != nil {
		return nil, apperror.AppError{
			BusinessError: err.Error(),
			UserError:     "Ошибка при сохранении новой задачи",
		}
	}

	return taskEntityToModel(taskEntity), nil

}

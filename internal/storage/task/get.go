package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"todo/internal/service/models"
	"todo/internal/storage/entity"
	apperror "todo/pkg/errors"
)

func (s *TaskStorage) GetTask(ctx context.Context, id int) (*models.Task, error) {
	s.l.Info("Получение задачи по ID:", id)

	query, args, err := squirrel.
		Select("id", "title", "description", "status", "created_at", "updated_at").
		From("tasks").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, apperror.AppError{
			BusinessError: err.Error(),
			UserError:     "Ошибка при построении запроса",
		}
	}

	var taskEntity entity.TaskEntity
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
			UserError:     "Задача не найдена",
		}
	}

	return taskEntityToModel(taskEntity), nil
}

package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"todo/internal/service/models"
	"todo/internal/storage/entity"
	apperror "todo/pkg/errors"
)

func (s *TaskStorage) GetTasks(ctx context.Context, limit, offset int, order string) ([]*models.Task, error) {
	s.l.Info("Получение списка задач с пагинацией", "limit", limit, "offset", offset, "order", order)

	query, args, err := squirrel.
		Select("id", "title", "description", "status", "created_at", "updated_at").
		From("tasks").
		OrderBy("created_at " + order).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, apperror.AppError{
			BusinessError: err.Error(),
			UserError:     "Ошибка при построении запроса списка задач",
		}
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, apperror.AppError{
			BusinessError: err.Error(),
			UserError:     "Ошибка при получении списка задач",
		}
	}
	defer rows.Close()

	var modelsTasks []*models.Task
	for rows.Next() {
		var taskEntity entity.TaskEntity
		if err := rows.Scan(
			&taskEntity.ID,
			&taskEntity.Title,
			&taskEntity.Description,
			&taskEntity.Status,
			&taskEntity.CreatedAt,
			&taskEntity.UpdatedAt,
		); err != nil {
			return nil, apperror.AppError{
				BusinessError: err.Error(),
				UserError:     "Ошибка при чтении данных задачи",
			}
		}
		modelsTasks = append(modelsTasks, taskEntityToModel(taskEntity))
	}

	return modelsTasks, nil
}

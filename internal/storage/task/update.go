package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"todo/internal/storage/entity"
	apperror "todo/pkg/errors"
)

func (s *TaskStorage) UpdateTask(ctx context.Context, id int, taskEntity entity.TaskEntity) error {
	s.l.Info("Обновление задачи:", id)

	query, args, err := squirrel.
		Update("tasks").
		Set("title", taskEntity.Title).
		Set("description", taskEntity.Description).
		Set("status", taskEntity.Status).
		Set("updated_at", squirrel.Expr("now()")).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return apperror.AppError{
			BusinessError: err.Error(),
			UserError:     "Ошибка при построении запроса обновления задачи",
		}
	}

	res, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return apperror.AppError{
			BusinessError: err.Error(),
			UserError:     "Ошибка при выполнении запроса обновления задачи",
		}
	}

	if res.RowsAffected() == 0 {
		return apperror.AppError{
			BusinessError: "Нет затронутых строк",
			UserError:     "Задача не найдена",
		}
	}

	return nil
}

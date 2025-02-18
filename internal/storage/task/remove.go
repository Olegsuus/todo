package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	apperror "todo/pkg/errors"
)

func (s *TaskStorage) RemoveTask(ctx context.Context, id int) error {
	s.l.Info("Удаление задачи:", id)

	query, args, err := squirrel.
		Delete("tasks").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return apperror.AppError{
			BusinessError: err.Error(),
			UserError:     "Ошибка при построении запроса удаления задачи",
		}
	}

	res, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return apperror.AppError{
			BusinessError: err.Error(),
			UserError:     "Ошибка при удалении задачи",
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

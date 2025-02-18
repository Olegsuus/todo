package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"todo/internal/service/models"
	"todo/internal/storage/entity"
)

type TaskStorage struct {
	db *pgxpool.Pool
	l  *slog.Logger
}

func NewTaskStorage(db *pgxpool.Pool, l *slog.Logger) *TaskStorage {
	return &TaskStorage{
		db: db,
		l:  l,
	}
}

func taskEntityToModel(entity entity.TaskEntity) *models.Task {
	return &models.Task{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Status:      entity.Status,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

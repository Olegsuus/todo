package service

import (
	"todo/internal/controllers/rest/handlers/domain"
	"todo/internal/service/models"
	"todo/internal/storage"
	"todo/internal/storage/entity"
)

type TaskService struct {
	taskStorage storage.TaskStorage
}

func NewTaskService(taskStorage storage.TaskStorage) *TaskService {
	return &TaskService{
		taskStorage: taskStorage,
	}
}

func modelsToDomain(modelsTask *models.Task) *domain.Task {
	return &domain.Task{
		ID:          modelsTask.ID,
		Title:       modelsTask.Title,
		Description: modelsTask.Description,
		Status:      modelsTask.Status,
		CreatedAt:   modelsTask.CreatedAt,
		UpdatedAt:   modelsTask.UpdatedAt,
	}
}

func modelsToEntity(modelsTask *models.Task) *entity.TaskEntity {
	return &entity.TaskEntity{
		ID:          modelsTask.ID,
		Title:       modelsTask.Title,
		Description: modelsTask.Description,
		Status:      modelsTask.Status,
		CreatedAt:   modelsTask.CreatedAt,
		UpdatedAt:   modelsTask.UpdatedAt,
	}
}

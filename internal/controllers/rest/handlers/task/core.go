package handler

import (
	"log/slog"
	"todo/internal/service"
)

type TaskHandler struct {
	taskService service.TaskService
	l           *slog.Logger
}

func NewTaskHandler(taskService service.TaskService, l *slog.Logger) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
		l:           l,
	}
}

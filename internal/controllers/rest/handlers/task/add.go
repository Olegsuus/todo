package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"todo/internal/controllers/rest/handlers/domain"
	apperror "todo/pkg/errors"
)

// @Summary  Добавить задачу
// @Description Добавить новую задачу с указанными данными
// @Tags   Tasks
// @Accept  json
// @Produce json
// @Param    task body domain.TaskDTO  true  "Данные задачи"
// @Success  201  {object}  domain.Task
// @Failure  400  {object}  apperror.ReqError
// @Failure  500  {object}  apperror.ReqError
// @Router  /tasks [post]
func (h *TaskHandler) AddTask(c *fiber.Ctx) error {
	var dto domain.TaskDTO

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": apperror.ReqError{
				Status: apperror.Decoding.Status,
				Text:   apperror.Database.Text,
			},
		})
	}

	id, err := h.taskService.AddTask(c.Context(), dto.Title, dto.Description, dto.Status)
	if err != nil {
		h.l.Debug("ошибка при записи новой задачи", slog.String("details", fmt.Sprintf("err: %s", err)))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": apperror.ReqError{
				Status: apperror.Database.Status,
				Text:   apperror.Database.Text,
			},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(domain.AddTaskRequest{Success: true, ID: id})
}

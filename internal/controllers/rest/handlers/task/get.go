package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"strconv"
	apperror "todo/pkg/errors"
)

// @Summary   Получить задачу по ID
// @Description Возвращает задачу по её ID
// @Tags   tasks
// @Accept  json
// @Produce  json
// @Param    id path int true "ID задачи"
// @Success  200 {object} domain.Task
// @Failure  400 {object} apperror.ReqError
// @Failure  404 {object} apperror.ReqError
// @Router  /tasks/{id} [get]
func (h *TaskHandler) GetTask(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": apperror.ReqError{
				Status: apperror.Decoding.Status,
				Text:   apperror.Database.Text,
			},
		})
	}

	task, err := h.taskService.GetTask(c.Context(), id)
	if err != nil {
		h.l.Debug("ошибка при получении задачи", slog.String("details", fmt.Sprintf("err: %s", err)))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": apperror.ReqError{
				Status: apperror.Database.Status,
				Text:   apperror.Database.Text,
			},
		})
	}
	return c.JSON(task)
}

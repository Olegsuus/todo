package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"strconv"
	"todo/internal/controllers/rest/handlers/domain"
	apperror "todo/pkg/errors"
)

// @Summary     Удалить задачу
// @Description Удаляет задачу по указанному ID
// @Tags       tasks
// @Accept     json
// @Produce    json
// @Param      id    path int true "ID задачи"
// @Success    204   {object} domain.SuccessRequest
// @Failure    400   {object}  apperror.ReqError
// @Failure    500   {object}  apperror.ReqError
// @Router  /tasks/{id} [delete]
func (h *TaskHandler) RemoveTask(c *fiber.Ctx) error {
	ctx := c.Context()
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": apperror.ReqError{
				Status: apperror.Decoding.Status,
				Text:   apperror.Decoding.Text,
			},
		})
	}

	if err = h.taskService.RemoveTask(ctx, id); err != nil {
		h.l.Debug("ошибка при удалении задачи", slog.String("details", fmt.Sprintf("err: %s", err)))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": apperror.ReqError{
				Status: apperror.Database.Status,
				Text:   apperror.Database.Text,
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.SuccessRequest{Success: true})
}

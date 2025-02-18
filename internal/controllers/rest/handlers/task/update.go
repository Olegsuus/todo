package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"strconv"
	"todo/internal/controllers/rest/handlers/domain"
	apperror "todo/pkg/errors"
)

// @Summary    Обновить задачу
// @Description Обновляет данные существующей задачи
// @Tags       Tasks
// @Accept     json
// @Produce    json
// @Param      id path int true "ID задачи"
// @Param      task  body      domain.TaskDTO true "Обновленные данные задачи"
// @Success    200   {object}  domain.SuccessRequest
// @Failure    400   {object}  apperror.ReqError
// @Failure    500   {object}  apperror.ReqError
// @Router  /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
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

	var dto domain.TaskDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": apperror.ReqError{
				Status: apperror.Decoding.Status,
				Text:   apperror.Database.Text,
			},
		})
	}

	if err = h.taskService.UpdateTask(ctx, id, dto.Title, dto.Description, dto.Status); err != nil {
		h.l.Debug("ошибка при обновлении задачи", slog.String("details", fmt.Sprintf("err: %s", err)))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": apperror.ReqError{
				Status: apperror.Database.Status,
				Text:   apperror.Database.Text,
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.SuccessRequest{Success: true})
}

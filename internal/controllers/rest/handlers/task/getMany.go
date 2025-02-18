package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	apperror "todo/pkg/errors"
)

// @Summary    Получить список задач
// @Description Возвращает список задач с поддержкой пагинации и сортировки
// @Tags      Tasks
// @Accept    json
// @Produce   json
// @Param     limit   query int false "Лимит"
// @Param     page  query int  false  "Страница"
// @Param     order   query string  false  "Порядок сортировки (asc, desc)"
// @Success   200     {array} domain.Task
// @Failure   500     {object} apperror.ReqError
// @Router   /tasks [get]
func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	ctx := c.Context()

	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 0)
	order := c.Query("order", "desc")

	tasks, err := h.taskService.GetTasks(ctx, limit, page, order)
	if err != nil {
		h.l.Debug("ошибка при получении списка задач", slog.String("details", fmt.Sprintf("err: %s", err)))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": apperror.ReqError{
				Status: apperror.Database.Status,
				Text:   apperror.Database.Text,
			},
		})
	}
	return c.JSON(tasks)
}

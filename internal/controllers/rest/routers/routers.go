package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "todo/docs"
	handler "todo/internal/controllers/rest/handlers/task"
)

func RegisterRoutes(app *fiber.App, taskHandler *handler.TaskHandler) *fiber.Router {
	api := app.Group("/tasks")
	api.Post("/", taskHandler.AddTask)
	api.Get("/", taskHandler.GetTasks)
	api.Get("/:id", taskHandler.GetTask)
	api.Put("/:id", taskHandler.UpdateTask)
	api.Delete("/:id", taskHandler.RemoveTask)

	api.Get("/swagger/*", swagger.HandlerDefault)

	return &api
}

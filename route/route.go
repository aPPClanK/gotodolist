package route

import (
	"github.com/aPPClanK/gotodolist/handler"
	"github.com/aPPClanK/gotodolist/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	api := app.Group("/api", middleware.JWTMiddleware)
	api.Get("/tasks", handler.GetTasks)
	api.Get("/tasks/:id", handler.GetTaskById)
	api.Post("/tasks/", handler.CreateTask)
	api.Patch("/tasks/:id", handler.UpdateTaskById)
	api.Delete("tasks/:id", handler.DeleteTaskById)
}

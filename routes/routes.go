package routes

import (
	"github.com/aPPClanK/gotodolist/handlers"
	"github.com/aPPClanK/gotodolist/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register", handlers.RegisterHandler)
	app.Post("/login", handlers.LoginHandler)
	app.Post("/logout", handlers.LogoutHandler)
	//app.Get("/users", handlers.GetUsersHandler)

	api := app.Group("/api", middleware.JWTMiddleware)
	api.Get("/tasks", handlers.GetTasksHandler)
	api.Get("/tasks/:id", handlers.GetTaskByIdHandler)
	api.Post("/tasks/", handlers.CreateTaskHandler)
	api.Patch("/tasks/:id", handlers.UpdateTaskByIdHandler)
	api.Delete("tasks/:id", handlers.DeleteTaskByIdHandler)
}

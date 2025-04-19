package handlers

import (
	"github.com/aPPClanK/gotodolist/db"
	"github.com/aPPClanK/gotodolist/models"

	"github.com/gofiber/fiber/v2"
)

func GetTasksHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var tasks []models.Task

	result := db.DB.Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Get tasks error"})
	}

	return c.JSON(tasks)
}
func GetTaskByIdHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var task models.Task

	result := db.DB.First(&task, "id = ? AND user_id = ?", id, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found or not yours"})
	}

	return c.JSON(task)
}
func CreateTaskHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if task.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title is required"})
	}

	task.UserID = userID
	if err := db.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Create failed"})
	}

	return c.JSON(task)
}

func UpdateTaskByIdHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var task models.Task

	result := db.DB.First(&task, "id = ? AND user_id = ?", id, userID)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found or not yours"})

	}

	task.Completed = !task.Completed
	if err := db.DB.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Update failed"})
	}

	return c.JSON(task)
}

func DeleteTaskByIdHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var task models.Task

	result := db.DB.First(&task, "id = ? AND user_id = ?", id, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found or not yours"})
	}

	if err := db.DB.Delete(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Delete failed"})
	}

	return c.JSON(fiber.Map{"status": "success"})
}

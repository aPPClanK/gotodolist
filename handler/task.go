package handler

import (
	"github.com/aPPClanK/gotodolist/database"
	"github.com/aPPClanK/gotodolist/model"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var tasks []model.Task

	result := database.DB.Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Get tasks error"})
	}

	return c.JSON(tasks)
}
func GetTaskById(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var task model.Task

	result := database.DB.First(&task, "id = ? AND user_id = ?", id, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found or not yours"})
	}

	return c.JSON(task)
}
func CreateTask(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var task model.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if task.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title is required"})
	}

	task.UserID = userID
	if err := database.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Create failed"})
	}

	return c.JSON(task)
}

func UpdateTaskById(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var task model.Task

	result := database.DB.First(&task, "id = ? AND user_id = ?", id, userID)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found or not yours"})

	}

	task.Completed = !task.Completed
	if err := database.DB.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Update failed"})
	}

	return c.JSON(task)
}

func DeleteTaskById(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id := c.Params("id")

	var task model.Task

	result := database.DB.First(&task, "id = ? AND user_id = ?", id, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found or not yours"})
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Delete failed"})
	}

	return c.JSON(fiber.Map{"status": "success"})
}

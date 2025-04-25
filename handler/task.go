package handler

import (
	"errors"

	"github.com/aPPClanK/gotodolist/database"
	"github.com/aPPClanK/gotodolist/model"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
)

func GetUserID(c *fiber.Ctx) (uint, error) {
	user := c.Locals("user")
	if user == nil {
		return 0, errors.New("user not found in context")
	}

	token, ok := user.(*jwt.Token)
	if !ok {
		return 0, errors.New("failed to assert user to jwt.Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("failed to assert claims to jwt.MapClaims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id is missing or invalid in token claims")
	}

	return uint(userID), nil
}

func GetTasks(c *fiber.Ctx) error {
	userID, err := GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	var tasks []model.Task

	result := database.DB.Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "get tasks error"})
	}

	return c.Status(fiber.StatusOK).JSON(tasks)
}

func GetTaskById(c *fiber.Ctx) error {
	userID, err := GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	id := c.Params("id")

	var task model.Task

	result := database.DB.First(&task, "id = ? AND user_id = ?", id, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found or not yours"})
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func CreateTask(c *fiber.Ctx) error {
	userID, err := GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	var task model.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	if task.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "title is required"})
	}

	task.UserID = userID
	if err := database.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "create failed"})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func UpdateTaskById(c *fiber.Ctx) error {
	userID, err := GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	id := c.Params("id")

	var task model.Task

	result := database.DB.First(&task, "id = ? AND user_id = ?", id, userID)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found or not yours"})

	}

	task.Completed = !task.Completed
	if err := database.DB.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "update failed"})
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func DeleteTaskById(c *fiber.Ctx) error {
	userID, err := GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	id := c.Params("id")

	var task model.Task

	result := database.DB.First(&task, "id = ? AND user_id = ?", id, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found or not yours"})
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "delete failed"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

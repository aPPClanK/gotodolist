package handlers

import (
	"github.com/aPPClanK/gotodolist/db"
	"github.com/aPPClanK/gotodolist/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsersHandler(c *fiber.Ctx) error {
	var users []models.User

	result := db.DB.Select("id", "name").Preload("Tasks").Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Get users error",
		})
	}

	return c.JSON(fiber.Map{"status": "success", "data": users})
}

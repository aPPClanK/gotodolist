package handlers

import (
	"time"

	"github.com/aPPClanK/gotodolist/db"
	"github.com/aPPClanK/gotodolist/models"
	"github.com/aPPClanK/gotodolist/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func RegisterHandler(c *fiber.Ctx) error {
	var input AuthInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Hash error"})
	}

	user := models.User{Name: input.Name, Password: hashedPassword}
	if err := db.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exists"})
	}

	return c.JSON(fiber.Map{"status": "Registration successful"})
}

func LoginHandler(c *fiber.Ctx) error {
	var input AuthInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := db.DB.Where("name = ?", input.Name).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token generation failed"})
	}

	// secure true и samesite "None" для связи фронта и бека, если фронт и бек на разных портах/доменах
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
	})

	return c.JSON(fiber.Map{"status": "Login successful"})
}

func LogoutHandler(c *fiber.Ctx) error {
	c.ClearCookie("jwt")
	return c.JSON(fiber.Map{"status": "Logout successful"})
}

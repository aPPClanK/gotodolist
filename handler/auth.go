package handler

import (
	"github.com/aPPClanK/gotodolist/database"
	"github.com/aPPClanK/gotodolist/middleware"
	"github.com/aPPClanK/gotodolist/model"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

type AuthInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(c *fiber.Ctx) error {
	var input AuthInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Hash error"})
	}

	user := model.User{Name: input.Name, Password: hashedPassword}
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exists"})
	}

	return c.JSON(fiber.Map{"status": "Registration successful"})
}

func Login(c *fiber.Ctx) error {
	var input AuthInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user model.User
	if err := database.DB.Where("name = ?", input.Name).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	if !CheckPasswordHash(input.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
	}

	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token generation failed"})
	}

	return c.JSON(token)
}

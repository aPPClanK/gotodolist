package main

import (
	"log"
	"os"

	"github.com/aPPClanK/gotodolist/database"
	"github.com/aPPClanK/gotodolist/route"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	if err := database.Connect(); err != nil {
		log.Fatal("DB connect error: ", err)
	}

	app := fiber.New()
	route.SetupRoutes(app)

	host := os.Getenv("APP_HOST")
	if host == "" {
		host = "localhost"
	}
	log.Fatal(app.Listen(host + ":3000"))
}

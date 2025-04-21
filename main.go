package main

import (
	"log"

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

	log.Fatal(app.Listen("localhost:3000"))
}

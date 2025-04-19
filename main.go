package main

import (
	"log"

	"github.com/aPPClanK/gotodolist/db"
	"github.com/aPPClanK/gotodolist/routes"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	if err := db.Connect(); err != nil {
		log.Fatal("DB connect error: ", err)
	}

	app := fiber.New()
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://localhost:5173",
	// 	AllowCredentials: true,
	// }))
	routes.SetupRoutes(app)

	log.Fatal(app.Listen("localhost:3000"))
}

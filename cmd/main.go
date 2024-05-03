package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kauefraga/inus/internal/services"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/users", services.CreateUser)

	app.Listen(":3000")
}

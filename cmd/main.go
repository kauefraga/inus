package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kauefraga/inus/internal/database"
	"github.com/kauefraga/inus/internal/services"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	db := database.Connect()
	defer db.Close()

	app.Post("/users", func(c *fiber.Ctx) error {
		return services.CreateUser(c, db)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return services.LoginUser(c, db)
	})

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("jwtsecretkey")},
	}))

	// DELETE "/users"

	app.Listen(":3000")
}

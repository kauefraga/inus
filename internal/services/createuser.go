package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kauefraga/inus/internal/domain"
)

func CreateUser(c *fiber.Ctx) error {
	user := domain.User{}
	err := c.BodyParser(&user)
	if err != nil {
		panic(err)
	}

	return c.Status(201).JSON(user)
}

package services

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func DeleteUser(c *fiber.Ctx, db *sql.DB) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["user"].(string)

	_, err := db.Exec("DELETE FROM users WHERE name = $1", name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "Failed deleting account.",
		})
	}

	return c.SendStatus(204)
}

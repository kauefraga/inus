package services

import (
	"database/sql"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kauefraga/inus/internal/validators"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginUser(c *fiber.Ctx, db *sql.DB) error {
	user := UserLoginRequest{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "Invalid body.",
		})
	}

	if validators.IsUserNameInvalid(user.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "Username is invalid. Username must have more than 3 and less than 51 characters.",
		})
	}

	user.Name = strings.ToLower(user.Name)

	var password string
	err = db.QueryRow("SELECT password FROM users WHERE name = $1", user.Name).Scan(&password)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(&fiber.Map{
			"error": "User does not exist.",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "The given password does not match with the actual user password.",
		})
	}

	claims := jwt.MapClaims{
		"authorized": true,
		"user":       user.Name,
		"exp":        time.Now().Add(time.Hour * 6),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("jwtsecretkey"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "Failed signing token.",
		})
	}

	// TODO: send a cookie instead of a response body
	return c.JSON(&fiber.Map{
		"token": tokenString,
	})
}

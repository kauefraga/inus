package services

import (
	"database/sql"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kauefraga/inus/internal/domain"
	"github.com/kauefraga/inus/internal/validators"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx, db *sql.DB) error {
	user := domain.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "Invalid body.",
		})
	}

	if validators.IsUserNameInvalid(user.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "Username invalid. Username must have more than 3 and less than 51 characters.",
		})
	}

	if validators.IsEmailInvalid(user.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "E-mail invalid. E-mail must not be empty.",
		})
	}

	user.Name = strings.ToLower(user.Name)

	var id int
	db.QueryRow("SELECT id FROM users WHERE name = $1", user.Name).Scan(&id)
	if id != 0 {
		return c.Status(fiber.StatusConflict).JSON(&fiber.Map{
			"error": "User already exists.",
		})
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": "Failed hashing password.",
		})
	}
	user.Password = string(hashed)

	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "Failed inserting user in the database.",
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
	return c.Status(201).JSON(&fiber.Map{
		"token": tokenString,
	})
}

package services

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kauefraga/inus/internal/database"
	"github.com/kauefraga/inus/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	user := domain.User{}

	err := c.BodyParser(&user)
	if err != nil {
		log.Println(c.Method(), c.Path(), err)
	}

	hasUsernameInvalidlength := len(user.Name) < 4 || len(user.Name) > 50
	isUsernameEmpty := len(strings.TrimSpace(user.Name)) == 0

	if hasUsernameInvalidlength || isUsernameEmpty {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "Username invalid. Username must have more than 3 and less than 51 characters.",
		})
	}

	// TODO: validate user e-mail

	user.Name = strings.ToLower(user.Name)

	_, ok := database.DB[user.Name]
	if ok {
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
	database.DB[user.Name] = user

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

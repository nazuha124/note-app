package routes

import (
	"os"
	"time"

	"notes-backend/database"
	"notes-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/")
	api.Post("/register", register)
	api.Post("/login", login)
}

type registerReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func register(c *fiber.Ctx) error {
	var body registerReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	// check existing
	var exists int
	err := database.DB.Get(&exists, "SELECT count(1) FROM users WHERE email=$1", body.Email)
	if err == nil && exists > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "email already used"})
	}
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "server error"})
	}
	res, err := database.DB.Exec("INSERT INTO users (email,password) VALUES ($1,$2)", body.Email, string(hash))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "db error",
			"debug": err.Error(), // tambahkan baris ini
		})
	}
	_ = res
	return c.JSON(fiber.Map{"message": "registered"})
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(c *fiber.Ctx) error {
	var body loginReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	var user models.User
	err := database.DB.Get(&user, "SELECT id,email,password FROM users WHERE email=$1", body.Email)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}
	// create token with user id
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not create token"})
	}
	return c.JSON(fiber.Map{"token": t})
}

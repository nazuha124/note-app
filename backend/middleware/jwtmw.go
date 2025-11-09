package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func RequireAuth() fiber.Handler {
	secret := os.Getenv("JWT_SECRET")
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(secret),
		ContextKey:   "user", // stores *jwt.Token into ctx.Locals("user")
		TokenLookup:  "header:Authorization",
		AuthScheme:   "Bearer",
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

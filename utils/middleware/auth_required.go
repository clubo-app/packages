package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func AuthRequired(secret string) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(secret),
		ErrorHandler: jwtError,
	})
}

func AuthOptional(secret string) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
		Filter: func(c *fiber.Ctx) bool {
			return c.Get("Authorization") == ""
		},
	})
}
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing or malformed Token")
	}
	return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired Token")
}

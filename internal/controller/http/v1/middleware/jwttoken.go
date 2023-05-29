package middleware

import (
	"clean-architecture-service/internal/tokens"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   tokens.JwtSignatureKey,
		ContextKey:   "jwt",
	})
}

func jwtError(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"StatusCode": fiber.StatusBadRequest,
			"Message":    "Missing or malformed JWT",
			"Data":       nil,
			"Error":      err,
		})
	} else {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"StatusCode": fiber.StatusUnauthorized,
			"Message":    "Invalid or expired JWT",
			"Data":       nil,
			"Error":      err,
		})
	}
}

package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GetUserIDAsStr(c *fiber.Ctx) (userID string) {
	jwtData := c.Locals("jwt").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	return id
}

func GetUserIDAsUUID(c *fiber.Ctx) (userID uuid.UUID, err error) {
	IDString := GetUserIDAsStr(c)
	UserID, err := uuid.Parse(IDString)
	if err != nil {
		return uuid.New(), err
	}
	return UserID, nil
}

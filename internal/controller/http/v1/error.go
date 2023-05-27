package v1

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorUserAlreadyExists = errors.New("user already exists")
	ErrorUserValidation    = errors.New("user validation failed")
)

func errorResponse(ctx *fiber.Ctx, code int, msg string, data interface{}, err error) error {
	return ctx.Status(code).JSON(Response{
		StatusCode: code,
		Message:    msg,
		Data:       data,
		Error:      err.Error(),
	})
}

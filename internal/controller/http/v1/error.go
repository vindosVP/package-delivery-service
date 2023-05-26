package v1

import "github.com/gofiber/fiber/v2"

func errorResponse(ctx *fiber.Ctx, code int, msg string, err error) error {
	return ctx.Status(code).JSON(Response{
		StatusCode: code,
		Message:    msg,
		Data:       nil,
		Error:      err,
	})
}

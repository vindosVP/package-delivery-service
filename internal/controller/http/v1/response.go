package v1

import "github.com/gofiber/fiber/v2"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

func OkResponse(ctx *fiber.Ctx, code int, msg string, data interface{}) error {
	return ctx.Status(code).JSON(Response{
		StatusCode: code,
		Message:    msg,
		Data:       data,
		Error:      nil,
	})
}

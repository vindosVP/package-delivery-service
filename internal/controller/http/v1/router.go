package v1

import (
	_ "clean-architecture-service/docs/swagger"
	"clean-architecture-service/internal/usecase"
	"clean-architecture-service/pkg/logger"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// swag init -g internal/controller/http/v1/router.go --output docs/swagger

// SetupRouter -.
// Swagger spec:
// @title       Delivery service API
// @description Delivery service
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func SetupRouter(handler *fiber.App, u usecase.User, p usecase.Package, d usecase.Delivery, l logger.Interface) {
	handler.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	handler.Get("/swagger/*", swagger.HandlerDefault)

	h := handler.Group("/v1")
	{
		SetUserRoutes(h, u, l)
		SetPackageRoutes(h, p, l)
		SetDeliveryRoutes(h, d, l)
	}
}

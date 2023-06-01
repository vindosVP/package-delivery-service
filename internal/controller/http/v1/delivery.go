package v1

import (
	"clean-architecture-service/internal/controller/http/v1/middleware"
	"clean-architecture-service/internal/usecase"
	"clean-architecture-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type DeliveryRoutes struct {
	d usecase.Delivery
	l logger.Interface
}

func SetDeliveryRoutes(handler fiber.Router, d usecase.Delivery, l logger.Interface) {
	r := &DeliveryRoutes{
		d: d,
		l: l,
	}
	h := handler.Group("/deliveries")
	h.Post("", middleware.Protected(), r.create)
}

func (r DeliveryRoutes) create(c *fiber.Ctx) error {
	panic(1)
}

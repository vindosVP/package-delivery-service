package v1

import (
	"clean-architecture-service/internal/usecase"
	"clean-architecture-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserRoutes struct {
	u usecase.User
	l logger.Interface
}

type registerUserRequest struct {
	Email           string `json:"email" binding:"required"  example:"vadiminmail@gmail.com"`
	Password        string `json:"password" binding:"required"  example:"qwerty123"`
	Name            string `json:"name" binding:"required"  example:"Vadim"`
	LastName        string `json:"lastName" binding:"required"  example:"Valov"`
	DeliveryAddress string `json:"deliveryAddress" binding:"required"  example:"Pushkina street"`
}

type registerUserResponse struct {
	ID              uuid.UUID `json:"id" binding:"required"  example:"d9e48656-ae36-4fde-af78-5f6250e11ead"`
	Email           string    `json:"email" binding:"required"  example:"vadiminmail@gmail.com"`
	Name            string    `json:"name" binding:"required"  example:"Vadim"`
	LastName        string    `json:"lastName" binding:"required"  example:"Valov"`
	DeliveryAddress string    `json:"deliveryAddress" binding:"required"  example:"Pushkina street"`
}

func SetUserRoutes(handler fiber.Router, u usecase.User, l logger.Interface) {
	r := &UserRoutes{
		u: u,
		l: l,
	}
	h := handler.Group("/user")
	h.Post("/register", r.register)
}

// @Summary     Register
// @Description Register a new user
// @ID          register
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param       request body registerUserRequest true "User data"
// @Success     200 {object} registerUserResponse
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /user/register [post]
func (r *UserRoutes) register(c *fiber.Ctx) error {
	user := &registerUserRequest{}
	if err := c.BodyParser(user); err != nil {
		r.l.Error(err, "http - v1 - c.BodyParser")
		return errorResponse(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	res, err := r.u.Register(
		user.Email,
		user.Password,
		user.Name,
		user.LastName,
		user.DeliveryAddress,
	)

	if err != nil {
		r.l.Error(err, "http - v1 - r.u.Register")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to create user", err)
	}

	responseData := &registerUserResponse{
		ID:              res.ID,
		Email:           res.Email,
		Name:            res.Name,
		LastName:        res.LastName,
		DeliveryAddress: res.DeliveryAddress,
	}

	return OkResponse(c, fiber.StatusOK, "User created", responseData)
}

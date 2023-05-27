package v1

import (
	"clean-architecture-service/internal/usecase"
	"clean-architecture-service/internal/validations"
	"clean-architecture-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoutes struct {
	u usecase.User
	l logger.Interface
}

type registerUserRequest struct {
	Email           string `json:"email" binding:"required"  example:"vadiminmail@gmail.com" validate:"required,email"`
	Password        string `json:"password" binding:"required"  example:"qwerty123" validate:"required,min=8,max=50"`
	Name            string `json:"name" binding:"required"  example:"Vadim" validate:"required"`
	LastName        string `json:"lastName" binding:"required"  example:"Valov" validate:"required"`
	DeliveryAddress string `json:"deliveryAddress" binding:"required"  example:"Pushkina street"  validate:"required"`
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
		r.l.Error(err, "v1 - register - c.BodyParser")
		return errorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil, err)
	}

	isValid, errs := validations.UniversalValidation(user)

	if !isValid {
		return errorResponse(c, fiber.StatusBadRequest, "Validation error", errs, ErrorUserValidation)
	}

	exists, err := r.u.UserExists(user.Email)

	if err != nil && err != gorm.ErrRecordNotFound {
		r.l.Error(err, "http - v1 - r.u.UserExists")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to check if user already exists", nil, err)
	}

	if exists {
		return errorResponse(c, fiber.StatusBadRequest, "User with this email already exists", nil, ErrorUserAlreadyExists)
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
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to create user", nil, err)
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

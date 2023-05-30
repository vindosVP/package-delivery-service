package v1

import (
	"clean-architecture-service/internal/controller/http/v1/middleware"
	"clean-architecture-service/internal/usecase"
	"clean-architecture-service/internal/validations"
	"clean-architecture-service/pkg/database"
	"clean-architecture-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserRoutes struct {
	u usecase.User
	l logger.Interface
}

type updateUserRequest struct {
	Password        string `json:"password" binding:"required"  example:"qwerty123" validate:"required,min=8,max=50"`
	Name            string `json:"name" binding:"required"  example:"Vadim" validate:"required"`
	LastName        string `json:"lastName" binding:"required"  example:"Valov" validate:"required"`
	DeliveryAddress string `json:"deliveryAddress" binding:"required"  example:"Pushkina street"  validate:"required"`
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

type updateUserResponse struct {
	ID              uuid.UUID `json:"id" binding:"required"  example:"d9e48656-ae36-4fde-af78-5f6250e11ead"`
	Email           string    `json:"email" binding:"required"  example:"vadiminmail@gmail.com"`
	Name            string    `json:"name" binding:"required"  example:"Vadim"`
	LastName        string    `json:"lastName" binding:"required"  example:"Valov"`
	DeliveryAddress string    `json:"deliveryAddress" binding:"required"  example:"Pushkina street"`
}

type authUserRequest struct {
	Email    string `json:"email" binding:"required"  example:"vadiminmail@gmail.com" validate:"required,email"`
	Password string `json:"password" binding:"required"  example:"qwerty123" validate:"required,min=8,max=50"`
}

type refreshAuthRequest struct {
	UserID       uuid.UUID `json:"user_id" validate:"required"`
	RefreshToken string    `json:"refresh_token" validate:"required"`
}

func SetUserRoutes(handler fiber.Router, u usecase.User, l logger.Interface) {
	r := &UserRoutes{
		u: u,
		l: l,
	}
	h := handler.Group("/users")
	h.Post("/user/register", r.register)
	h.Post("/user/auth", r.auth)
	h.Post("/user/refresh", middleware.GormTransaction(database.DB, l), r.refresh)
	h.Patch("/user/update", middleware.Protected(), r.update)
}

// @Summary     Update
// @Description Updates users data
// @ID          update
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param       request body updateUserRequest true "User data"
// @Success     200 {object} Response
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /users/user/update [patch]
func (r *UserRoutes) update(c *fiber.Ctx) error {
	req := &updateUserRequest{}
	if err := c.BodyParser(req); err != nil {
		r.l.Error(err, "v1 - register - c.BodyParser")
		return errorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil, err)
	}
	isValid, errs := validations.UniversalValidation(req)

	if !isValid {
		return errorResponse(c, fiber.StatusBadRequest, "Validation error", errs, ErrorValidationFailed)
	}

	id := GetUserIDFromJWT(c)

	userID, err := uuid.Parse(id)
	if err != nil {
		r.l.Error(err, "v1 - register - uuid.Parse")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to parse the uuid", nil, err)
	}

	user, err := r.u.Update(
		userID,
		req.Password,
		req.Name,
		req.LastName,
		req.DeliveryAddress,
	)

	if err != nil {
		if err == usecase.ErrorUserDoesNotExist {
			return errorResponse(c, fiber.StatusBadRequest, "User with this id does not exist", nil, err)
		} else {
			r.l.Error(err, "http - v1 - r.u.Update")
			return errorResponse(c, fiber.StatusInternalServerError, "Failed to update user", nil, err)
		}
	}

	res := updateUserResponse{
		ID:              user.ID,
		Email:           user.Email,
		Name:            user.Name,
		LastName:        user.LastName,
		DeliveryAddress: user.DeliveryAddress,
	}

	return OkResponse(c, fiber.StatusOK, "User updated successfully", res)
}

// @Summary     Refresh
// @Description Refreshes users JWT token
// @ID          refresh
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param       request body refreshAuthRequest true "User data"
// @Success     200 {object} Response
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /users/user/refresh [post]
func (r *UserRoutes) refresh(c *fiber.Ctx) error {
	req := &refreshAuthRequest{}
	if err := c.BodyParser(req); err != nil {
		r.l.Error(err, "v1 - register - c.BodyParser")
		return errorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil, err)
	}

	isValid, errs := validations.UniversalValidation(req)

	if !isValid {
		return errorResponse(c, fiber.StatusBadRequest, "Validation error", errs, ErrorValidationFailed)
	}

	res, err := r.u.Refresh(req.RefreshToken, req.UserID)
	if err != nil {
		if err == usecase.ErrorInvalidToken {
			return errorResponse(c, fiber.StatusBadRequest, "Invalid token", nil, err)
		} else {
			r.l.Error(err, "http - v1 - r.u.Auth")
			return errorResponse(c, fiber.StatusInternalServerError, "Failed to refresh token", nil, err)
		}
	}

	return OkResponse(c, fiber.StatusOK, "Refreshed successfully", res)
}

// @Summary     Auth
// @Description Authenticates user
// @ID          auth
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param       request body authUserRequest true "User data"
// @Success     200 {object} Response
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /users/user/auth [post]
func (r *UserRoutes) auth(c *fiber.Ctx) error {
	req := &authUserRequest{}
	if err := c.BodyParser(req); err != nil {
		r.l.Error(err, "v1 - register - c.BodyParser")
		return errorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil, err)
	}

	isValid, errs := validations.UniversalValidation(req)

	if !isValid {
		return errorResponse(c, fiber.StatusBadRequest, "Validation error", errs, ErrorValidationFailed)
	}

	res, err := r.u.Auth(req.Email, req.Password)
	if err != nil {
		if err == usecase.ErrorInvalidEmailOrPwd {
			return errorResponse(c, fiber.StatusBadRequest, "Invalid email or password", nil, err)
		} else {
			r.l.Error(err, "http - v1 - r.u.Auth")
			return errorResponse(c, fiber.StatusInternalServerError, "Failed to autenticate user", nil, err)
		}
	}

	return OkResponse(c, fiber.StatusOK, "Authenticated successfully", res)
}

// @Summary     Register
// @Description Register a new user
// @ID          register
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param       request body registerUserRequest true "User data"
// @Success     201 {object} registerUserResponse
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /users/user/register [post]
func (r *UserRoutes) register(c *fiber.Ctx) error {
	req := &registerUserRequest{}
	if err := c.BodyParser(req); err != nil {
		r.l.Error(err, "v1 - register - c.BodyParser")
		return errorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil, err)
	}

	isValid, errs := validations.UniversalValidation(req)

	if !isValid {
		return errorResponse(c, fiber.StatusBadRequest, "Validation error", errs, ErrorValidationFailed)
	}

	res, err := r.u.Register(
		req.Email,
		req.Password,
		req.Name,
		req.LastName,
		req.DeliveryAddress,
	)

	if err != nil {
		if err == usecase.ErrorUserAlreadyExists {
			return errorResponse(c, fiber.StatusBadRequest, "User already exists", nil, err)
		} else {
			r.l.Error(err, "http - v1 - r.u.Register")
			return errorResponse(c, fiber.StatusInternalServerError, "Failed to create user", nil, err)
		}
	}

	responseData := &registerUserResponse{
		ID:              res.ID,
		Email:           res.Email,
		Name:            res.Name,
		LastName:        res.LastName,
		DeliveryAddress: res.DeliveryAddress,
	}

	return OkResponse(c, fiber.StatusCreated, "User created", responseData)
}

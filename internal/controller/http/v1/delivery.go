package v1

import (
	"clean-architecture-service/internal/controller/http/v1/middleware"
	"clean-architecture-service/internal/usecase"
	"clean-architecture-service/internal/validations"
	"clean-architecture-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DeliveryRoutes struct {
	d usecase.Delivery
	l logger.Interface
}

var statuses = [3]string{"new", "in progress", "delivered"}

type deliveryRequest struct {
	RecipientID uuid.UUID `json:"recipientID" binding:"required"  example:"1873eecd-c2d0-4aa2-a8d4-e0de232c5ac6" validate:"required"`
	PackageID   uuid.UUID `json:"packageID" binding:"required"  example:"9ebb7be6-c8fc-49a5-b941-f8090c0db7fc" validate:"required"`
	Urgent      bool      `json:"urgent" binding:"required"  example:"true"`
}

type setStatusRequest struct {
	Status string `json:"status" binding:"required" example:"delivered" validate:"required"`
}

type deliveryResponse struct {
	DeliveryID  uuid.UUID `json:"DeliveryID" example:"1873eecd-c2d0-4aa2-a8d4-e0de232c5ac6"`
	SenderId    uuid.UUID `json:"SenderId" example:"9ebb7be6-c8fc-49a5-b941-f8090c0db7fc"`
	RecipientID uuid.UUID `json:"recipientID" example:"9ebb7be6-c8fc-49a5-b941-f8090c0db7fc"`
	PackageID   uuid.UUID `json:"packageID" example:"9ebb7be6-c8fc-49a5-b941-f8090c0db7fc"`
	Urgent      bool      `json:"urgent" example:"false"`
	Status      string    `json:"status" example:"new"`
}

func SetDeliveryRoutes(handler fiber.Router, d usecase.Delivery, l logger.Interface) {
	r := &DeliveryRoutes{
		d: d,
		l: l,
	}
	h := handler.Group("/deliveries")
	h.Post("", middleware.Protected(), r.create)
	h.Get("", middleware.Protected(), r.getDeliveries)
	h.Get("/:DeliveryID", middleware.Protected(), r.getDeliveries)
	h.Post("/:DeliveryID/setStatus", middleware.Protected(), r.setStatus)
	h.Patch("/:DeliveryID", middleware.Protected(), r.updateDelivery)
}

// @Summary     Update delivery
// @Description Updates delivery by id
// @ID          updateDelivery
// @Tags  	    deliveries
// @Accept      json
// @Produce     json
// @Param 		deliveryID path string true "delivery ID" example(6155c774-d1e2-4816-b7f4-52ebb949f044)
// @Param       request body deliveryRequest true "New status"
// @Success     200 {object} Response{data=deliveryResponse}
// @Failure     500 {object} Response
// @Router      /deliveries/{deliveryID} [patch]
func (r DeliveryRoutes) updateDelivery(c *fiber.Ctx) error {
	req := &deliveryRequest{}
	if err := c.BodyParser(req); err != nil {
		r.l.Error(err, "v1 - Update - c.BodyParser")
		return errorResponse(c, fiber.StatusBadRequest, MsgInvalidRequestBody, nil, err)
	}
	isValid, errs := validations.UniversalValidation(req)

	if !isValid {
		return errorResponse(c, fiber.StatusBadRequest, MsgNotValid, errs, ErrorValidationFailed)
	}

	deliveryIDStr := c.Params("DeliveryID")
	deliveryID, err := uuid.Parse(deliveryIDStr)
	if err != nil {
		r.l.Error(err, "v1 - setStatus - uuid.Parse")
		return errorResponse(c, fiber.StatusInternalServerError, MsgFailedToParseUUID, nil, err)
	}

	del, err := r.d.Update(deliveryID, req.RecipientID, req.PackageID, req.Urgent)
	if err != nil {
		if err == usecase.ErrorCantSendToYourself {
			return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil, err)
		}
		if err == usecase.ErrorPackageDoesNotExist {
			return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil, err)
		}
		if err == usecase.ErrorRecipientDoesNotExist {
			return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil, err)
		}
		r.l.Error(err, "v1 - Update - r.d.Update")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to update delivery", nil, err)
	}

	res := deliveryResponse{
		DeliveryID:  del.ID,
		SenderId:    del.SenderID,
		RecipientID: del.RecipientID,
		PackageID:   del.PackageID,
		Urgent:      del.Urgent,
		Status:      del.Status,
	}

	return OkResponse(c, fiber.StatusOK, "Delivery updated successfully", res)
}

// @Summary     Get delivery
// @Description Returns delivery by id
// @ID          getDelivery
// @Tags  	    deliveries
// @Accept      json
// @Produce     json
// @Param 		deliveryID path string true "delivery ID" example(6155c774-d1e2-4816-b7f4-52ebb949f044)
// @Success     200 {object} Response{data=deliveryResponse}
// @Failure     500 {object} Response
// @Router      /deliveries/{deliveryID} [get]
func (r DeliveryRoutes) getDelivery(c *fiber.Ctx) error {
	deliveryIDStr := c.Params("DeliveryID")
	deliveryID, err := uuid.Parse(deliveryIDStr)
	if err != nil {
		r.l.Error(err, "v1 - getDelivery - uuid.Parse")
		return errorResponse(c, fiber.StatusInternalServerError, MsgFailedToParseUUID, nil, err)
	}

	del, err := r.d.GetDelivery(deliveryID)
	if err != nil {
		if err == usecase.ErrorDeliveryDoesNotExist {
			return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil, err)
		}
		r.l.Error(err, "v1 - getDelivery - r.d.GetDelivery")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to get delivery", nil, err)
	}

	res := deliveryResponse{
		DeliveryID:  del.ID,
		SenderId:    del.SenderID,
		RecipientID: del.RecipientID,
		PackageID:   del.PackageID,
		Urgent:      del.Urgent,
		Status:      del.Status,
	}

	return OkResponse(c, fiber.StatusOK, "", res)
}

// @Summary     Get deliveries
// @Description Returns all deliveries
// @ID          getDeliveries
// @Tags  	    deliveries
// @Accept      json
// @Produce     json
// @Success     200 {object} Response{data=[]deliveryResponse}
// @Failure     500 {object} Response
// @Router      /deliveries [get]
func (r DeliveryRoutes) getDeliveries(c *fiber.Ctx) error {
	deliveries, err := r.d.Get()
	if err != nil {
		r.l.Error(err, "v1 - getDeliveries - r.d.Get")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to get deliveries", nil, err)
	}

	res := make([]deliveryResponse, len(deliveries), cap(deliveries))
	for i, v := range deliveries {

		res[i] = deliveryResponse{
			DeliveryID:  v.ID,
			SenderId:    v.SenderID,
			RecipientID: v.RecipientID,
			PackageID:   v.PackageID,
			Urgent:      v.Urgent,
			Status:      v.Status,
		}

	}

	return OkResponse(c, fiber.StatusOK, "", res)
}

// @Summary     Set status
// @Description Sets status to delivery
// @ID          setStatus
// @Tags  	    deliveries
// @Accept      json
// @Produce     json
// @Param 		deliveryID path string true "delivery ID" example(6155c774-d1e2-4816-b7f4-52ebb949f044)
// @Param       request body setStatusRequest true "New status"
// @Success     200 {object} Response
// @Failure     500 {object} Response
// @Router      /deliveries/{deliveryID}/setStatus [post]
func (r DeliveryRoutes) setStatus(c *fiber.Ctx) error {
	req := &setStatusRequest{}
	if err := c.BodyParser(req); err != nil {
		r.l.Error(err, "v1 - setStatus - c.BodyParser")
		return errorResponse(c, fiber.StatusBadRequest, MsgInvalidRequestBody, nil, err)
	}
	isValid, errs := validations.UniversalValidation(req)

	if !isValid {
		return errorResponse(c, fiber.StatusBadRequest, MsgNotValid, errs, ErrorValidationFailed)
	}

	if !IsValidStatus(req.Status) {
		return errorResponse(c, fiber.StatusBadRequest, "Not valid status", nil, ErrorValidationFailed)
	}

	deliveryIDStr := c.Params("DeliveryID")
	deliveryID, err := uuid.Parse(deliveryIDStr)
	if err != nil {
		r.l.Error(err, "v1 - setStatus - uuid.Parse")
		return errorResponse(c, fiber.StatusInternalServerError, MsgFailedToParseUUID, nil, err)
	}

	if err := r.d.SetStatus(deliveryID, req.Status); err != nil {
		r.l.Error(err, "v1 - setStatus - r.d.SetStatus")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to set status", nil, err)
	}

	return OkResponse(c, fiber.StatusOK, "Status set successfully", nil)
}

// @Summary     Create delivery
// @Description Creates delivery
// @ID          create
// @Tags  	    deliveries
// @Accept      json
// @Produce     json
// @Param       request body deliveryRequest true "Delivery data"
// @Success     200 {object} Response{data=deliveryResponse}
// @Failure     500 {object} Response
// @Router      /deliveries [post]
func (r DeliveryRoutes) create(c *fiber.Ctx) error {
	req := &deliveryRequest{}
	if err := c.BodyParser(req); err != nil {
		r.l.Error(err, "v1 - create - c.BodyParser")
		return errorResponse(c, fiber.StatusBadRequest, MsgInvalidRequestBody, nil, err)
	}
	isValid, errs := validations.UniversalValidation(req)

	if !isValid {
		return errorResponse(c, fiber.StatusBadRequest, MsgNotValid, errs, ErrorValidationFailed)
	}

	UserID, err := GetUserIDAsUUID(c)
	if err != nil {
		r.l.Error(err, "v1 - create - GetUserIDAsUUID")
		return errorResponse(c, fiber.StatusInternalServerError, MsgFailedToParseUUID, nil, err)
	}

	del, err := r.d.Create(
		UserID,
		req.RecipientID,
		req.PackageID,
		req.Urgent,
		"new",
	)

	if err != nil {
		if err == usecase.ErrorCantSendToYourself {
			return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil, err)
		}
		if err == usecase.ErrorSenderDoesNotExist {
			return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil, err)
		}
		if err == usecase.ErrorPackageDoesNotExist {
			return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil, err)
		}
		if err == usecase.ErrorRecipientDoesNotExist {
			return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil, err)
		}
		r.l.Error(err, "v1 - create - r.d.Create")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to crate delivery", nil, err)
	}

	res := deliveryResponse{
		DeliveryID:  del.ID,
		SenderId:    del.SenderID,
		RecipientID: del.RecipientID,
		PackageID:   del.PackageID,
		Urgent:      del.Urgent,
		Status:      del.Status,
	}

	return OkResponse(c, fiber.StatusCreated, "Delivery created successfully", res)
}

func IsValidStatus(status string) bool {
	for _, v := range statuses {
		if v == status {
			return true
		}
	}

	return false
}

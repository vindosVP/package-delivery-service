package v1

import (
	"clean-architecture-service/internal/controller/http/v1/middleware"
	"clean-architecture-service/internal/usecase"
	"clean-architecture-service/internal/validations"
	"clean-architecture-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PackageRoutes struct {
	p usecase.Package
	l logger.Interface
}

func SetPackageRoutes(handler fiber.Router, p usecase.Package, l logger.Interface) {
	r := &PackageRoutes{
		p: p,
		l: l,
	}
	h := handler.Group("/packages")
	h.Post("", middleware.Protected(), r.create)
	h.Get("", middleware.Protected(), r.getPackages)
	h.Patch("/:packageID", middleware.Protected(), r.update)
}

type createPackageRequest struct {
	Name   string  `json:"name" binding:"required"  example:"Package for Moxem"`
	Weight float64 `json:"weight" binding:"required"  example:"11.3"`
	Height float64 `json:"height" binding:"required"  example:"15"`
	Width  float64 `json:"width" binding:"required"  example:"13.8"`
}

type createPackageResponse struct {
	PackageID uuid.UUID `json:"packageID" example:"6155c774-d1e2-4816-b7f4-52ebb949f044"`
	OwnerID   uuid.UUID `json:"ownerID" example:"P1873eecd-c2d0-4aa2-a8d4-e0de232c5ac6"`
	Name      string    `json:"name" example:"Package for Moxem"`
	Weight    float64   `json:"weight" example:"11.3"`
	Height    float64   `json:"height" example:"15"`
	Width     float64   `json:"width" example:"13.8"`
}

type getPackagesResponse struct {
	Packages []createPackageResponse `json:"packages"`
}

// @Summary     Get packages
// @Description Returns user`s packages
// @ID          getPackages
// @Tags  	    packages
// @Accept      json
// @Produce     json
// @Success     200 {object} Response
// @Failure     500 {object} Response
// @Router      /packages [get]
func (r *PackageRoutes) getPackages(c *fiber.Ctx) error {
	id := GetUserIDFromJWT(c)
	UserID, err := uuid.Parse(id)
	if err != nil {
		r.l.Error(err, "v1 - getPackages - uuid.Parse")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to parse the uuid", nil, err)
	}

	packages, err := r.p.GetPackages(UserID)
	if err != nil {
		r.l.Error(err, "v1 - getPackages - r.p.GetPackages")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to get packages", nil, err)
	}

	pkgs := make([]createPackageResponse, len(packages), len(packages))

	for i, v := range packages {
		pack := createPackageResponse{
			PackageID: v.ID,
			OwnerID:   v.OwnerID,
			Name:      v.Name,
			Weight:    v.Weight,
			Height:    v.Height,
			Width:     v.Width,
		}
		pkgs[i] = pack
	}

	res := getPackagesResponse{
		Packages: pkgs,
	}

	return OkResponse(c, fiber.StatusOK, "OK", res)
}

// @Summary     Update
// @Description Updates a package
// @ID          updatePackage
// @Tags  	    packages
// @Accept      json
// @Produce     json
// @Param       request body createPackageRequest true "User data"
// @Success     200 {object} Response
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /packages [patch]
func (r *PackageRoutes) update(c *fiber.Ctx) error {
	req := &createPackageRequest{}
	if err := c.BodyParser(req); err != nil {
		r.l.Error(err, "v1 - update - c.BodyParser")
		return errorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil, err)
	}
	isValid, errs := validations.UniversalValidation(req)

	if !isValid {
		return errorResponse(c, fiber.StatusBadRequest, "Validation error", errs, ErrorValidationFailed)
	}

	packageIDStr := c.Params("packageID")
	packageID, err := uuid.Parse(packageIDStr)
	if err != nil {
		r.l.Error(err, "v1 - update - uuid.Parse")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to parse the uuid", nil, err)
	}

	id := GetUserIDFromJWT(c)
	UserID, err := uuid.Parse(id)
	if err != nil {
		r.l.Error(err, "v1 - update - uuid.Parse")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to parse the uuid", nil, err)
	}

	pack, err := r.p.Update(
		UserID,
		packageID,
		req.Name,
		req.Weight,
		req.Height,
		req.Width,
	)

	if err != nil {
		if err == usecase.ErrorPackageDoesNotExist {
			return errorResponse(c, fiber.StatusBadRequest, "Package does not exist", nil, err)
		}
		if err == usecase.ErrorPackageDoesNotBelongToUser {
			return errorResponse(c, fiber.StatusBadRequest, "Package does not belong to this user", nil, err)
		}
		r.l.Error(err, "v1 - update - r.p.Update")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed update package", nil, err)
	}

	res := createPackageResponse{
		PackageID: pack.ID,
		OwnerID:   pack.OwnerID,
		Name:      pack.Name,
		Weight:    pack.Weight,
		Height:    pack.Height,
		Width:     pack.Width,
	}

	return OkResponse(c, fiber.StatusOK, "Package updated", res)
}

// @Summary     Create
// @Description Crates a package
// @ID          create
// @Tags  	    packages
// @Accept      json
// @Produce     json
// @Param       request body createPackageRequest true "User data"
// @Success     201 {object} Response
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /packages [post]
func (r *PackageRoutes) create(c *fiber.Ctx) error {
	req := &createPackageRequest{}
	if err := c.BodyParser(req); err != nil {
		r.l.Error(err, "v1 - create - c.BodyParser")
		return errorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil, err)
	}
	isValid, errs := validations.UniversalValidation(req)

	if !isValid {
		return errorResponse(c, fiber.StatusBadRequest, "Validation error", errs, ErrorValidationFailed)
	}

	id := GetUserIDFromJWT(c)
	ownerID, err := uuid.Parse(id)
	if err != nil {
		r.l.Error(err, "v1 - create - uuid.Parse")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to parse the uuid", nil, err)
	}

	pack, err := r.p.Create(
		ownerID,
		req.Name,
		req.Weight,
		req.Height,
		req.Width,
	)

	if err != nil {
		r.l.Error(err, "v1 - create - r.p.Create")
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to crate package", nil, err)
	}

	res := createPackageResponse{
		PackageID: pack.ID,
		OwnerID:   pack.OwnerID,
		Name:      pack.Name,
		Weight:    pack.Weight,
		Height:    pack.Height,
		Width:     pack.Width,
	}

	return OkResponse(c, fiber.StatusCreated, "Package created", res)
}

package usecase

import (
	"clean-architecture-service/internal/entity"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PackageUseCase struct {
	packageRepo PackageRepo
}

func NewPackageUseCase(pr PackageRepo) *PackageUseCase {
	return &PackageUseCase{
		packageRepo: pr,
	}
}

func (p PackageUseCase) GetPackage(ownerID uuid.UUID, packageID uuid.UUID) (*entity.Package, error) {
	pack, err := p.packageRepo.FindByID(packageID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrorPackageDoesNotExist
		}
		return nil, fmt.Errorf("PackageUsecase - GetPackage - p.packageRepo.FindByID")
	}

	if pack.OwnerID != ownerID {
		return nil, ErrorPackageDoesNotBelongToUser
	}

	return pack, nil
}

func (p PackageUseCase) GetPackages(ownerID uuid.UUID) ([]entity.Package, error) {
	packages, err := p.packageRepo.GetPackages(ownerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []entity.Package{}, nil
		}
		return []entity.Package{}, err
	}

	return packages, nil
}

func (p PackageUseCase) Create(ownerID uuid.UUID, name string, weight float64, height float64, width float64, status string) (*entity.Package, error) {
	newPack, err := p.packageRepo.Create(&entity.Package{
		OwnerID: ownerID,
		Name:    name,
		Weight:  weight,
		Height:  height,
		Width:   width,
		Status:  status,
	})
	if err != nil {
		return nil, fmt.Errorf("PackageUsecase - Create - p.packageRepo.Create")
	}

	return newPack, nil
}

func (p PackageUseCase) Update(userID uuid.UUID, packageID uuid.UUID, name string, weight float64, height float64, width float64) (*entity.Package, error) {
	pack, err := p.packageRepo.FindByID(packageID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrorPackageDoesNotExist
		} else {
			return nil, fmt.Errorf("PackageUsecase - Update - p.packageRepo.FindByID")
		}
	}

	if userID != pack.OwnerID {
		return nil, ErrorPackageDoesNotBelongToUser
	}

	crPack := &entity.Package{
		ID:      pack.ID,
		OwnerID: pack.OwnerID,
		Name:    name,
		Weight:  weight,
		Height:  height,
		Width:   width,
	}
	newPack, err := p.packageRepo.Update(crPack)

	return newPack, err
}

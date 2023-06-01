package usecase

import (
	"clean-architecture-service/internal/entity"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeliveryUseCase struct {
	DeliveryRepo DeliveryRepo
	UserRepo     UserRepo
	PackageRepo  PackageRepo
}

func NewDeliveryUseCase(delRepo DeliveryRepo, uRepo UserRepo, packRepo PackageRepo) *DeliveryUseCase {
	return &DeliveryUseCase{
		DeliveryRepo: delRepo,
		UserRepo:     uRepo,
		PackageRepo:  packRepo,
	}
}

func (d DeliveryUseCase) Create(senderID uuid.UUID, recipientID uuid.UUID, packageID uuid.UUID, urgent bool, status string) (*entity.Delivery, error) {
	packExists, err := d.PackageRepo.PackageExistsByID(packageID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("DeliveryUsecase - Create - d.PackageRepo.PackageExistsByID")
	}

	if !packExists {
		return nil, ErrorPackageDoesNotExist
	}

	senderExists, err := d.UserRepo.UserExistsByID(senderID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("DeliveryUsecase - Create - d.UserRepo.UserExistsByID(senderID)")
	}

	if !senderExists {
		return nil, ErrorSenderDoesNotExist
	}

	recipExists, err := d.UserRepo.UserExistsByID(recipientID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("DeliveryUsecase - Create - d.UserRepo.UserExistsByID(recipientID)")
	}

	if !recipExists {
		return nil, ErrorRecipientDoesNotExist
	}

	del := &entity.Delivery{
		SenderID:    senderID,
		RecipientID: recipientID,
		PackageID:   packageID,
		Urgent:      urgent,
		Status:      status,
	}

	newDel, err := d.DeliveryRepo.Create(del)
	if err != nil {
		return nil, fmt.Errorf("DeliveryUsecase - Create - d.DeliveryRepo.Create")
	}

	return newDel, nil
}

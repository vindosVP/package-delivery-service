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

func (d DeliveryUseCase) Update(deliveryID uuid.UUID, recipientID uuid.UUID, packageID uuid.UUID, urgent bool) (*entity.Delivery, error) {
	packExists, err := d.PackageRepo.PackageExistsByID(packageID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("DeliveryUsecase - Update - d.PackageRepo.PackageExistsByID")
	}

	if !packExists {
		return nil, ErrorPackageDoesNotExist
	}

	recipExists, err := d.UserRepo.UserExistsByID(recipientID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("DeliveryUsecase - Update - d.UserRepo.UserExistsByID(recipientID)")
	}

	if !recipExists {
		return nil, ErrorRecipientDoesNotExist
	}

	del, err := d.DeliveryRepo.FindByID(deliveryID)
	if err != nil {
		return nil, fmt.Errorf("DeliveryUsecase - Update - d.DeliveryRepo.FindByID")
	}

	if del.SenderID == recipientID {
		return nil, ErrorCantSendToYourself
	}

	updDel := &entity.Delivery{
		ID:          del.ID,
		SenderID:    del.SenderID,
		RecipientID: recipientID,
		PackageID:   packageID,
		Urgent:      urgent,
		Status:      del.Status,
	}

	newDel, err := d.DeliveryRepo.Update(updDel)
	if err != nil {
		return nil, fmt.Errorf("DeliveryUsecase - Update - d.DeliveryRepo.Update")
	}

	return newDel, nil
}

func (d DeliveryUseCase) GetDelivery(deliveryID uuid.UUID) (*entity.Delivery, error) {
	del, err := d.DeliveryRepo.FindByID(deliveryID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrorDeliveryDoesNotExist
		}
		return nil, fmt.Errorf("DeliveryUsecase - GetDelivery - d.DeliveryRepo.FindByID")
	}
	return del, nil
}

func (d DeliveryUseCase) Get() ([]entity.Delivery, error) {
	deliveries, err := d.DeliveryRepo.Get()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return make([]entity.Delivery, 0, 0), nil
		}
		return nil, fmt.Errorf("DeliveryUsecase - Get - d.DeliveryRepo.Get")
	}

	return deliveries, nil
}

func (d DeliveryUseCase) SetStatus(deliveryID uuid.UUID, newStatus string) error {
	delExist, err := d.DeliveryRepo.DeliveryExists(deliveryID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("DeliveryUsecase - SetStatus - d.DeliveryRepo.DeliveryExists")
	}

	if !delExist {
		return ErrorDeliveryDoesNotExist
	}

	if err := d.DeliveryRepo.SetStatus(deliveryID, newStatus); err != nil {
		return err
	}

	return nil
}

func (d DeliveryUseCase) Create(senderID uuid.UUID, recipientID uuid.UUID, packageID uuid.UUID, urgent bool, status string) (*entity.Delivery, error) {
	if senderID == recipientID {
		return nil, ErrorCantSendToYourself
	}

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

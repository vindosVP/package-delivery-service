package delivery_repo

import (
	"clean-architecture-service/internal/entity"
	"gorm.io/gorm"
)

type DeliveryRepo struct {
	db *gorm.DB
}

func NewDeliveryRepo(db *gorm.DB) *DeliveryRepo {
	return &DeliveryRepo{
		db: db,
	}
}

func (d DeliveryRepo) Create(delivery *entity.Delivery) (*entity.Delivery, error) {
	return delivery, d.db.Create(&delivery).Error
}

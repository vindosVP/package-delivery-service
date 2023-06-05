package delivery_repo

import (
	"clean-architecture-service/internal/entity"
	"github.com/google/uuid"
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

func (d DeliveryRepo) DeliveryExists(deliveryID uuid.UUID) (bool, error) {
	var count int64
	err := d.db.Model(&entity.Delivery{}).Where("id = ?", deliveryID).Count(&count).Error
	return count > 0, err
}

func (d DeliveryRepo) FindByID(deliveryID uuid.UUID) (*entity.Delivery, error) {
	var delivery entity.Delivery
	err := d.db.Where("id = ?", deliveryID).First(&delivery).Error
	return &delivery, err
}

func (d DeliveryRepo) SetStatus(deliveryID uuid.UUID, newStatus string) error {
	return d.db.Model(&entity.Delivery{}).Where("id = ?", deliveryID).Update("status", newStatus).Error
}

func (d DeliveryRepo) Get() ([]entity.Delivery, error) {
	var deliveries []entity.Delivery
	err := d.db.Model(&entity.Delivery{}).Find(&deliveries).Error
	return deliveries, err
}

func (d DeliveryRepo) Update(delivery *entity.Delivery) (*entity.Delivery, error) {
	var newDel entity.Delivery
	err := d.db.Model(entity.Delivery{}).Where("id = ?", delivery.ID).Updates(delivery).First(&newDel).Error
	return &newDel, err
}

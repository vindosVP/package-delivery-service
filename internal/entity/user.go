package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;primary_key;unique;not_null;default:uuid_generate_v4()" json:"id"`
	Email           string    `gorm:"type:varchar(255);not_null;unique" json:"email"`
	Password        string    `gorm:"not_null"`
	Name            string    `gorm:"type:varchar(255);not_null" json:"name"`
	LastName        string    `gorm:"type:varchar(255);not_null" json:"lastName"`
	DeliveryAddress string    `gorm:"type:varchar(255);not_null" json:"deliveryAddress"`
}

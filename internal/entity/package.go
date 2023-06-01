package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Package struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key;unique;not_null;default:uuid_generate_v4()" json:"ID"`
	Owner   User      `gorm:"foreignkey:OwnerID"`
	OwnerID uuid.UUID `gorm:"type:uuid;not_null" json:"ownerID"`
	Name    string    `gorm:"type:varchar(255);not_null" json:"name"`
	Status  string    `gorm:"type:varchar(255);not_null" json:"status"`
	Weight  float64   `gorm:"not_null" json:"weight"`
	Height  float64   `gorm:"not_null" json:"height"`
	Width   float64   `gorm:"not_null" json:"width"`
}

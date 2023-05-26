package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Package struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key;unique;not_null" json:"ID"`
	Owner   User      `gorm:"foreignkey:OwnerID"`
	OwnerID uuid.UUID `gorm:"type:uuid;not_null" json:"ownerID"`
	Type    string    `gorm:"type:varchar(255);not_null" json:"type"`
	Weight  int       `gorm:"not_null" json:"weight"`
	Height  int       `gorm:"not_null" json:"height"`
	Width   int       `gorm:"not_null" json:"width"`
}

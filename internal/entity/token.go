package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	ID      int       `gorm:"primary_key;unique;not_null;auto_increment" json:"ID"`
	User    User      `gorm:"foreignkey:UserID"`
	UserID  uuid.UUID `gorm:"type:uuid;not_null" json:"userID"`
	Token   string    `gorm:"type:varchar(255);not_null" json:"token"`
	Revoked bool      `gorm:"not_null" json:"revoked"`
}

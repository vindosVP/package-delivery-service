package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Delivery struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;unique;not_null;default:uuid_generate_v4()" json:"ID"`
	SenderID    uuid.UUID `gorm:"type:uuid;not_null" json:"senderID"`
	Sender      User      `gorm:"foreignkey:SenderID"`
	RecipientID uuid.UUID `gorm:"type:uuid;not_null" json:"recipientID"`
	Recipient   User      `gorm:"foreignkey:RecipientID"`
	PackageID   uuid.UUID `gorm:"type:uuid;not_null" json:"packageID"`
	Package     Package   `gorm:"foreignkey:PackageID"`
	Urgent      bool      `gorm:"not_null" json:"urgent"`
	Delivered   bool      `gorm:"not_null" json:"delivered"`
}

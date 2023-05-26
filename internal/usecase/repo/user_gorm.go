package repo

import (
	"clean-architecture-service/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(user *entity.User) (*entity.User, error) {
	return user, u.db.Create(&user).Error
}

package user_repo

import (
	"clean-architecture-service/internal/entity"
	"github.com/google/uuid"
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

func (u *UserRepository) UserExists(email string) (bool, error) {
	var count int64
	err := u.db.Model(&entity.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (u *UserRepository) UserExistsByID(userID uuid.UUID) (bool, error) {
	var count int64
	err := u.db.Model(&entity.User{}).Where("id = ?", userID).Count(&count).Error
	return count > 0, err
}

func (u *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *UserRepository) FindByID(UserID uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("id = ?", UserID).First(&user).Error
	return &user, err
}

func (u *UserRepository) Update(user *entity.User) (*entity.User, error) {
	var newUser entity.User
	err := u.db.Model(&entity.User{}).Where("id = ?", user.ID).Updates(user).First(&newUser).Error
	return &newUser, err
}

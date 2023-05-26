package usecase

import (
	"clean-architecture-service/internal/entity"
	"clean-architecture-service/pkg/utils"
	"fmt"
)

type UserUseCase struct {
	userRepo UserRepo
}

func New(ur UserRepo) *UserUseCase {
	return &UserUseCase{
		userRepo: ur,
	}
}

func (u UserUseCase) Register(email string, password string, name string, lastName string, delAddr string) (*entity.User, error) {

	hashPwd, err := utils.GeneratePassword(password)

	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Register - utils.GeneratePassword")
	}

	return u.userRepo.Create(&entity.User{
		Email:           email,
		Password:        hashPwd,
		Name:            name,
		LastName:        lastName,
		DeliveryAddress: delAddr,
	})

}

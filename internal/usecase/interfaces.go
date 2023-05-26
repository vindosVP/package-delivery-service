package usecase

import (
	"clean-architecture-service/internal/entity"
)

type User interface {
	Register(email string, password string, name string, lastName string, delAddr string) (*entity.User, error)
}

type UserRepo interface {
	Create(user *entity.User) (*entity.User, error)
	//UserExists(email string) (bool, error)
}

package usecase

import (
	"clean-architecture-service/internal/entity"
	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type User interface {
	Register(email string, password string, name string, lastName string, delAddr string) (*entity.User, error)
	Auth(email string, password string) (map[string]interface{}, error)
	Refresh(token string, UserID uuid.UUID) (map[string]interface{}, error)
	Update(userID uuid.UUID, password string, name string, lastName string, delAddr string) (*entity.User, error)
}

type UserRepo interface {
	Create(user *entity.User) (*entity.User, error)
	UserExists(email string) (bool, error)
	UserExistsByID(userID uuid.UUID) (bool, error)
	FindByEmail(email string) (*entity.User, error)
	FindByID(UserID uuid.UUID) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
}

type TokenRepo interface {
	Create(userID uuid.UUID, token string) error
	GetActiveToken(UserID uuid.UUID) (*entity.Token, error)
	Revoke(UserID uuid.UUID) error
}

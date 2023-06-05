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

type Package interface {
	Create(ownerID uuid.UUID, name string, weight float64, height float64, width float64, status string) (*entity.Package, error)
	Update(userID uuid.UUID, packageID uuid.UUID, name string, weight float64, height float64, width float64) (*entity.Package, error)
	GetPackages(ownerID uuid.UUID) ([]entity.Package, error)
	GetPackage(UserID uuid.UUID, packageID uuid.UUID) (*entity.Package, error)
}

type Delivery interface {
	Create(senderID uuid.UUID, recipientID uuid.UUID, packageID uuid.UUID, urgent bool, status string) (*entity.Delivery, error)
	SetStatus(deliveryID uuid.UUID, newStatus string) error
	Get() ([]entity.Delivery, error)
	GetDelivery(deliveryID uuid.UUID) (*entity.Delivery, error)
	Update(deliveryID uuid.UUID, recipientID uuid.UUID, packageID uuid.UUID, urgent bool) (*entity.Delivery, error)
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

type PackageRepo interface {
	Create(pack *entity.Package) (*entity.Package, error)
	FindByID(packageID uuid.UUID) (*entity.Package, error)
	Update(pack *entity.Package) (*entity.Package, error)
	GetPackages(ownerID uuid.UUID) ([]entity.Package, error)
	PackageExistsByID(packageID uuid.UUID) (bool, error)
}

type DeliveryRepo interface {
	Create(delivery *entity.Delivery) (*entity.Delivery, error)
	DeliveryExists(deliveryID uuid.UUID) (bool, error)
	SetStatus(deliveryID uuid.UUID, newStatus string) error
	Get() ([]entity.Delivery, error)
	FindByID(deliveryID uuid.UUID) (*entity.Delivery, error)
	Update(delivery *entity.Delivery) (*entity.Delivery, error)
}

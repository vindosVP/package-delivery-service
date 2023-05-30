package usecase

import (
	"clean-architecture-service/internal/entity"
	"clean-architecture-service/internal/tokens"
	"clean-architecture-service/pkg/utils"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserUseCase struct {
	userRepo  UserRepo
	tokenRepo TokenRepo
}

func NewUserUseCase(ur UserRepo, tr TokenRepo) *UserUseCase {
	return &UserUseCase{
		userRepo:  ur,
		tokenRepo: tr,
	}
}

func (u UserUseCase) Update(userID uuid.UUID, password string, name string, lastName string, delAddr string) (*entity.User, error) {
	hashPwd, err := utils.GeneratePassword(password)

	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Update - utils.GeneratePassword")
	}

	exists, err := u.userRepo.UserExistsByID(userID)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("UserUsecase - Update - u.UserExists")
	}

	if !exists {
		return nil, ErrorUserDoesNotExist
	}

	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Update - u.userRepo.FindByID")
	}

	userUpd := &entity.User{
		ID:              userID,
		Password:        hashPwd,
		Email:           user.Email,
		Name:            name,
		LastName:        lastName,
		DeliveryAddress: delAddr,
	}

	newUser, err := u.userRepo.Update(userUpd)
	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Update - u.userRepo.Update")
	}

	return newUser, nil
}

func (u UserUseCase) Register(email string, password string, name string, lastName string, delAddr string) (*entity.User, error) {
	hashPwd, err := utils.GeneratePassword(password)

	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Register - utils.GeneratePassword")
	}

	exists, err := u.userRepo.UserExists(email)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("UserUsecase - Register - u.UserExists")
	}

	if exists {
		return nil, ErrorUserAlreadyExists
	}

	return u.userRepo.Create(&entity.User{
		Email:           email,
		Password:        hashPwd,
		Name:            name,
		LastName:        lastName,
		DeliveryAddress: delAddr,
	})
}

func (u UserUseCase) Refresh(token string, UserID uuid.UUID) (map[string]interface{}, error) {
	activeToken, err := u.tokenRepo.GetActiveToken(UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("UserUsecase - Refresh - u.tokenRepo.GetActiveToken")
	}

	if activeToken.Token != token {
		return nil, ErrorInvalidToken
	}

	err = u.tokenRepo.Revoke(UserID)
	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Refresh - u.tokenRepo.Revoke")
	}

	rt, err := tokens.GenerateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Refresh - tokens.GenerateRefreshToken")
	}

	err = u.tokenRepo.Create(UserID, rt)
	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Refresh - u.tokenRepo.Create")
	}

	user, err := u.userRepo.FindByID(UserID)
	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Refresh - u.userRepo.FindByID")
	}

	jwtToken, exp, err := tokens.GenerateJWT(user)
	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Refresh - tokens.GenerateJWT")
	}

	return map[string]interface{}{
		"user_id":       user.ID,
		"tokens":        jwtToken,
		"exp":           exp,
		"refresh_token": rt,
	}, nil
}

func (u UserUseCase) Auth(email string, password string) (map[string]interface{}, error) {
	exists, err := u.userRepo.UserExists(email)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("UserUsecase - Auth - u.UserExists")
	}

	if !exists {
		return nil, ErrorInvalidEmailOrPwd
	}

	user, err := u.userRepo.FindByEmail(email)

	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Auth - u.userRepo.FindByEmail")
	}

	if !utils.ComparePassword(user.Password, password) {
		return nil, ErrorInvalidEmailOrPwd
	}

	jwtToken, exp, err := tokens.GenerateJWT(user)
	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Auth - utils.GenerateJWT")
	}

	activeToken, err := u.tokenRepo.GetActiveToken(user.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("UserUsecase - Auth - u.tokenRepo.GetActiveToken")
	}

	if !activeToken.Revoked {
		err := u.tokenRepo.Revoke(user.ID)
		if err != nil {
			return nil, fmt.Errorf("UserUsecase - Auth - u.tokenRepo.Revoke")
		}
	}

	rt, err := tokens.GenerateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Auth - utils.GenerateRefreshToken")
	}

	err = u.tokenRepo.Create(user.ID, rt)
	if err != nil {
		return nil, fmt.Errorf("UserUsecase - Auth - u.tokenRepo.Create")
	}

	return map[string]interface{}{
		"user_id":       user.ID,
		"tokens":        jwtToken,
		"exp":           exp,
		"refresh_token": rt,
	}, nil
}

func (u UserUseCase) UserExists(email string) (bool, error) {
	return u.userRepo.UserExists(email)
}

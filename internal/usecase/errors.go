package usecase

import "errors"

var (
	ErrorUserAlreadyExists = errors.New("user already exists")
	ErrorInvalidEmailOrPwd = errors.New("invalid email or password")
	ErrorInvalidToken      = errors.New("invalid tokens")
	ErrorUserDoesNotExist  = errors.New("user with this id does not exist")
)

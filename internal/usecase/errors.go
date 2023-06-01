package usecase

import "errors"

var (
	ErrorUserAlreadyExists          = errors.New("user already exists")
	ErrorInvalidEmailOrPwd          = errors.New("invalid email or password")
	ErrorInvalidToken               = errors.New("invalid tokens")
	ErrorUserDoesNotExist           = errors.New("user with this id does not exist")
	ErrorPackageDoesNotExist        = errors.New("package with this id does not exist")
	ErrorPackageDoesNotBelongToUser = errors.New("package with this id does not belong to this user")
	ErrorSenderDoesNotExist         = errors.New("user with this id does not exist (sender)")
	ErrorRecipientDoesNotExist      = errors.New("user with this id does not exist (recipient)")
)

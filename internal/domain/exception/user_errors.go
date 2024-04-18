package exception

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user alredy exists")
	ErrUserNotFound      = errors.New("user not found")
)

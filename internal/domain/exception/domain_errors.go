package exception

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

var (
	ErrFolderAlreadyExists = errors.New("folder already exists")
	ErrFolderNotFound      = errors.New("folder not found")
)

var (
	ErrNameRequired = errors.New("name is required")
)

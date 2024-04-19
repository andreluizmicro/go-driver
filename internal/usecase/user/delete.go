package user

import (
	"errors"
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
	"github.com/andreluizmicro/go-driver-api/internal/domain/exception"
)

type DeleteUser struct {
	userRepository contracts.UserRepositoryInterface
}

func NewDeleteUser(userRepository contracts.UserRepositoryInterface) *DeleteUser {
	return &DeleteUser{
		userRepository: userRepository,
	}
}

func (us *DeleteUser) Execute(input DeleteInput) DeleteOutput {
	err := us.userRepository.Delete(input.ID)
	if err != nil && errors.Is(err, exception.ErrUserNotFound) {
		return DeleteOutput{
			Err: err,
		}
	}

	return DeleteOutput{
		Err: nil,
	}
}

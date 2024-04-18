package user

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
)

type DeleteUser struct {
	userRepository contracts.UserRepositoryInterface
}

func NewUserDelete(userRepository contracts.UserRepositoryInterface) *DeleteUser {
	return &DeleteUser{
		userRepository: userRepository,
	}
}

func (us *DeleteUser) Execute(input DeleteInput) DeleteOutput {
	err := us.userRepository.Delete(input.ID)
	return DeleteOutput{Success: err == nil}
}

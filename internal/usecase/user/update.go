package user

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
)

type UpdateUser struct {
	userRepository contracts.UserRepositoryInterface
}

func NewUpdateUser(userRepository contracts.UserRepositoryInterface) *UpdateUser {
	return &UpdateUser{
		userRepository: userRepository,
	}
}

func (u *UpdateUser) Execute(input UpdateInput) (*UpdateOutput, error) {
	user, err := u.userRepository.FindById(input.ID)
	if err != nil {
		return nil, err
	}
	err = user.Update(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = u.userRepository.Update(user)
	if err != nil {
		return nil, err
	}
	return &UpdateOutput{Success: true}, nil
}

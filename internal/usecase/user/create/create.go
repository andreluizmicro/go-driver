package create

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
)

type CreateUser struct {
	UserRepository contracts.UserRepositoryInterface
}

func NewCreateUser(userRepository contracts.UserRepositoryInterface) *CreateUser {
	return &CreateUser{
		UserRepository: userRepository,
	}
}

func (u *CreateUser) Execute(input Input) (*Output, error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	id, err := u.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return &Output{
		ID: *id,
	}, nil
}

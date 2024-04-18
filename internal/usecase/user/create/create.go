package create

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
)

type User struct {
	userRepository contracts.UserRepositoryInterface
}

func NewCreateUser(userRepository contracts.UserRepositoryInterface) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (u *User) Execute(input Input) (*Output, error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	id, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return &Output{
		ID: *id,
	}, nil
}

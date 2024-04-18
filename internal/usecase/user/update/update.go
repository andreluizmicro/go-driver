package update

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
)

type User struct {
	userRepository contracts.UserRepositoryInterface
}

func NewUpdateUser(userRepository contracts.UserRepositoryInterface) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (u *User) Execute(input Input) (*Output, error) {
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
	return &Output{Success: true}, nil
}

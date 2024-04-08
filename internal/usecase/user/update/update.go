package user

import "github.com/andreluizmicro/go-driver-api/internal/domain/contracts"

type UpdateUser struct {
	UserRepository contracts.UserRepositoryInterface
}

func NewUpdateUser(userRepository contracts.UserRepositoryInterface) *UpdateUser {
	return &UpdateUser{
		UserRepository: userRepository,
	}
}

func (u *UpdateUser) Execute(input Input) (*Output, error) {
	user, err := u.UserRepository.FindById(input.ID)
	if err != nil {
		return nil, err
	}
	user.Update(input.Name, input.Email, input.Password)

	err = u.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}
	return &Output{Success: true}, nil
}

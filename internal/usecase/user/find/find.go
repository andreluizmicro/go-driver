package find

import "github.com/andreluizmicro/go-driver-api/internal/domain/contracts"

type FindUser struct {
	userRepository contracts.UserRepositoryInterface
}

func NewFindUser(userRepository contracts.UserRepositoryInterface) *FindUser {
	return &FindUser{
		userRepository: userRepository,
	}
}

func (u *FindUser) Execute(input Input) (*Output, error) {
	user, err := u.userRepository.FindById(input.ID)
	if err != nil {
		return nil, err
	}

	return &Output{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		Password:   user.Password,
		ModifiedAt: user.ModifiedAt,
		LastLogin:  user.LastLogin,
	}, nil
}

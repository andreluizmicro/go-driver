package delete

import "github.com/andreluizmicro/go-driver-api/internal/domain/contracts"

type User struct {
	userRepository contracts.UserRepositoryInterface
}

func NewUserDelete(userRepository contracts.UserRepositoryInterface) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (us *User) Execute(input Input) error {
	return us.userRepository.Delete(input.ID)
}

package user

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository/filter"
)

type ListUser struct {
	userRepository contracts.UserRepositoryInterface
}

func NewListUser(userRepository contracts.UserRepositoryInterface) *ListUser {
	return &ListUser{
		userRepository: userRepository,
	}
}

func (u *ListUser) Execute(input ListInput) (*ListOutput, error) {
	filters := filter.NewFilters(
		input.Page,
		input.PerPage,
		input.Fields,
		input.Order,
		input.Email,
	)

	pagination, err := u.userRepository.FindAll(filters)
	if err != nil {
		return nil, err
	}

	return &ListOutput{
		Data:        pagination.GetItems(),
		Total:       pagination.GetTotal(),
		CurrentPage: pagination.GetCurrentPage(),
		TotalPage:   pagination.GetTotalPage(),
		FirstPage:   pagination.GetFirstPage(),
		LastPage:    pagination.GetLastPage(),
	}, nil
}

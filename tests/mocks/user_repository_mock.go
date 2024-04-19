package mocks

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository/filter"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *entity.User) (*int64, error) {
	args := m.Called(user)
	return args.Get(0).(*int64), args.Error(1)
}

func (m *UserRepositoryMock) FindAll(filters *filter.Filters) (contracts.PaginateInterface, error) {
	//TODO implement me
	panic("implement me")
}

func (m *UserRepositoryMock) FindById(id int64) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *UserRepositoryMock) Update(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (m *UserRepositoryMock) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (m *UserRepositoryMock) ExistsByEmail(s string) bool {
	//TODO implement me
	panic("implement me")
}

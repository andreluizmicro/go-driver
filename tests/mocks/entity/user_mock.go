package entity

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func (m *UserMock) NewUser() error {
	return entity.ErrEmailRequired
}

package contracts

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository/filter"
)

type UserRepositoryInterface interface {
	Create(*entity.User) (*int64, error)
	FindAll(*filter.Filters) (PaginateInterface, error)
	FindById(id int64) (*entity.User, error)
	Update(*entity.User) error
	Delete(id int64) error
	ExistsByEmail(string) bool
}

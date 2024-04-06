package contracts

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository"
)

type UserRepositoryInterface interface {
	Create(*entity.User) (*int64, error)
	FindAll(*repository.Filters) ([]entity.User, error)
	FindById(id int64) (*entity.User, error)
	Update(*entity.User) error
	Delete(id int64) error
	ExistsByEmail(string) bool
}

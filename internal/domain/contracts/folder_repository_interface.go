package contracts

import "github.com/andreluizmicro/go-driver-api/internal/domain/entity"

type FolderRepositoryInterface interface {
	Create(folder *entity.Folder) (int64, error)
	FindById(id int64) (*entity.Folder, error)
	Update(*entity.Folder) (bool, error)
	Delete(id int64) error
}

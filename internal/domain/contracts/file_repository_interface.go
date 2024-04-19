package contracts

import "github.com/andreluizmicro/go-driver-api/internal/domain/entity"

type FileRepositoryInterface interface {
	ListAll(folderId int64) ([]entity.File, error)
	Update(file entity.File) error
}

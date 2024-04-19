package folder

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
)

type CreateFolder struct {
	folderRepository contracts.FolderRepositoryInterface
}

func NewCreateFolder(folderRepository contracts.FolderRepositoryInterface) *CreateFolder {
	return &CreateFolder{
		folderRepository: folderRepository,
	}
}

func (u *CreateFolder) Execute(input CreateInput) (CreateOutput, error) {
	folder, err := entity.NewFolder(input.Name, input.ParentId)
	if err != nil {
		return CreateOutput{}, err
	}

	id, err := u.folderRepository.Create(folder)
	if err != nil {
		return CreateOutput{}, err
	}

	return CreateOutput{
		Id: id,
	}, nil
}

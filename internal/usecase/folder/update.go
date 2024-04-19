package folder

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
)

type UpdateFolder struct {
	folderRepository contracts.FolderRepositoryInterface
}

func NewUpdateFolder(folderRepository contracts.FolderRepositoryInterface) *UpdateFolder {
	return &UpdateFolder{
		folderRepository: folderRepository,
	}
}

func (u *UpdateFolder) Execute(input UpdateInput) (*UpdateOutput, error) {
	folder, err := u.folderRepository.FindById(input.ID)
	if err != nil {
		return nil, err
	}
	folder.Name = input.Name

	isUpdated, err := u.folderRepository.Update(folder)
	if err != nil {
		return &UpdateOutput{}, err
	}

	return &UpdateOutput{
		Success: isUpdated,
	}, nil
}

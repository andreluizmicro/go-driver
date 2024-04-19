package folder

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
)

type DeleteFolder struct {
	folderRepository contracts.FolderRepositoryInterface
	fileRepository   contracts.FileRepositoryInterface
}

func NewDeleteFolder(folderRepository contracts.FolderRepositoryInterface, fileRepository contracts.FileRepositoryInterface) *DeleteFolder {
	return &DeleteFolder{
		folderRepository: folderRepository,
		fileRepository:   fileRepository,
	}
}

func (u *DeleteFolder) Execute(input DeleteInput) (*DeleteOutput, error) {
	err := u.deleteFiles(input.ID)
	if err != nil {
		return nil, err
	}

	// TODO: listar folders

	err = u.folderRepository.Delete(input.ID)
	if err != nil {
		return &DeleteOutput{
			Success: false,
		}, err
	}

	return &DeleteOutput{
		Success: true,
	}, nil
}

func (u *DeleteFolder) deleteFiles(id int64) error {
	// tenta recuperar todos os arquivos de uma pasta
	files, err := u.fileRepository.ListAll(id)
	if err != nil {
		return err
	}

	// Tenta Deletar todos os files
	removedFiles := make([]entity.File, 0, len(files))
	for _, file := range files {
		file.Deleted = true
		err := u.fileRepository.Update(file)
		if err != nil {
			break
		}
		removedFiles = append(removedFiles, file)
	}

	if len(files) != len(removedFiles) {
		for _, file := range removedFiles {
			file.Deleted = false
			err := u.fileRepository.Update(file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

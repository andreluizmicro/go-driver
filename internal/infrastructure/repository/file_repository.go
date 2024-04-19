package repository

import (
	"database/sql"
	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
)

type FileRepository struct {
	db *sql.DB
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) ListAll(folderId int64) ([]entity.File, error) {
	stmt := `SELECT * FROM files where folder_id = ?`
	rows, err := r.db.Query(stmt, folderId)
	if err != nil {
		return nil, err
	}

	files := make([]entity.File, 0)
	for rows.Next() {
		var file entity.File
		err := rows.Scan(
			&file.ID,
			&file.FolderID,
			&file.OwnerID,
			&file.Name,
			&file.Type,
			&file.Path,
			&file.CreatedAt,
			&file.ModifiedAt,
			&file.Deleted,
		)
		if err != nil {
			continue
		}
		files = append(files, file)
	}

	return files, nil
}

func (r *FileRepository) Update(file entity.File) error {
	stmt := `UPDATE files SET name = ?, deleted = ? WHERE id = ?`
	_, err := r.db.Exec(stmt, file.Name, file.Deleted, file.ID)
	if err != nil {
		return err
	}
	return nil
}

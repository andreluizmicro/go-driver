package repository

import (
	"database/sql"
	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
	"github.com/andreluizmicro/go-driver-api/internal/domain/exception"
)

type FolderRepository struct {
	db *sql.DB
}

func NewFolderRepository(db *sql.DB) *FolderRepository {
	return &FolderRepository{
		db: db,
	}
}

func (r *FolderRepository) Create(folder *entity.Folder) (int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO folders (name, parent_id) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	res, err := stmt.Exec(folder.Name, folder.ParentID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *FolderRepository) FindById(id int64) (*entity.Folder, error) {
	var folder entity.Folder
	stmt := `SELECT * FROM folders WHERE id = ? AND deleted = ?`
	err := r.db.QueryRow(stmt, id, 0).Scan(
		&folder.ID,
		&folder.ParentID,
		&folder.Name,
		&folder.CreatedAt,
		&folder.ModifiedAt,
		&folder.Deleted,
	)
	if err != nil {
		return nil, exception.ErrFolderNotFound
	}

	return &folder, nil
}

func (r *FolderRepository) Update(folder *entity.Folder) (bool, error) {
	stmt, err := r.db.Prepare("UPDATE folders SET name = ? WHERE id = ?")
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(folder.Name, folder.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

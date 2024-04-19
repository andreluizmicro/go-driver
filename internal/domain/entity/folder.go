package entity

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/exception"
	"time"
)

type Folder struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	ParentID   int64     `json:"parent_id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
}

func NewFolder(name string, parentId int64) (*Folder, error) {
	folder := Folder{
		Name:     name,
		ParentID: parentId,
	}
	err := folder.Validate()
	if err != nil {
		return nil, err
	}
	return &folder, nil
}

func (f *Folder) Validate() error {
	if f.Name == "" {
		return exception.ErrNameRequired
	}
	return nil
}

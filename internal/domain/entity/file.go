package entity

import (
	"errors"
	"github.com/andreluizmicro/go-driver-api/internal/domain/exception"
	"time"
)

var (
	ErrOwnerRequired = errors.New("owner is required")
	ErrTypeRequired  = errors.New("type is required")
	ErrPathRequired  = errors.New("path is required")
)

type File struct {
	ID         int64     `json:"id"`
	FolderID   int64     `json:"-"`
	OwnerID    int64     `json:"Owner_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Path       string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
}

func (f *File) Validate() error {
	if f.OwnerID == 0 {
		return ErrOwnerRequired
	}
	if f.Name == "" {
		return exception.ErrNameRequired
	}
	if f.Type == "" {
		return ErrTypeRequired
	}
	if f.Path == "" {
		return ErrPathRequired
	}
	return nil
}

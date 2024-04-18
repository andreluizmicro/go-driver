package user

import (
	_ "github.com/go-playground/validator/v10"
	"time"
)

type CreateInput struct {
	Name     string `form:"name"`
	Email    string `form:"email" binding:"required,min=5"`
	Password string `form:"password"`
}

type CreateOutput struct {
	ID int64 `json:"id"`
}

type FindInput struct {
	ID int64 `uri:"id" binding:"required"`
}

type FindOutput struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	LastLogin  time.Time `json:"last_login"`
}

type UpdateInput struct {
	ID    int64  `uri:"id" binding:"required"`
	Name  string `form:"name"`
	Email string `form:"email" binding:"omitempty,required,min=5"`
}

type UpdateOutput struct {
	Success bool `json:"success"`
}

type DeleteInput struct {
	ID int64 `uri:"id" binding:"required"`
}

type DeleteOutput struct {
	Success bool `json:"success"`
}

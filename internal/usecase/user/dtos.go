package user

import (
	_ "github.com/go-playground/validator/v10"
	"time"
)

type CreateInput struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`
	Email    string `form:"email" binding:"required,min=5"`
	Password string `form:"password" binding:"required,min=8"`
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

type ListInput struct {
	Page    int64    `form:"page"`
	PerPage int64    `form:"per_page"`
	Fields  []string `form:"fields[]"`
	Order   string   `form:"order"`
	Email   string   `form:"email"`
}

type ListOutput struct {
	Data        []any `json:"data"`
	Total       int64 `json:"total"`
	CurrentPage int64 `json:"current_page"`
	TotalPage   int64 `json:"total_page"`
	FirstPage   int64 `json:"first_page"`
	LastPage    int64 `json:"last_page"`
}

type UpdateInput struct {
	ID       int64  `uri:"id" binding:"required"`
	Name     string `form:"name" binding:"required,min=5"`
	Email    string `form:"email" binding:"required,min=5"`
	Password string `form:"password" binding:"required,min=8"`
}

type UpdateOutput struct {
	Success bool `json:"success"`
}

type DeleteInput struct {
	ID int64 `uri:"id" binding:"required"`
}

type DeleteOutput struct {
	Err error
}

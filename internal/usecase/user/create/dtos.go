package create

import _ "github.com/go-playground/validator/v10"

type Input struct {
	Name     string `form:"name"`
	Email    string `form:"email" binding:"required,min=5"`
	Password string `form:"password"`
}

type Output struct {
	ID int64 `json:"id"`
}

package update

type Input struct {
	ID    int64  `uri:"id" binding:"required"`
	Name  string `form:"name"`
	Email string `form:"email" binding:"omitempty,required,min=5"`
}

type Output struct {
	Success bool `json:"success"`
}

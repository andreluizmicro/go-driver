package update

type Input struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" validate:"required,min=3,max=100"`
	Email string `json:"email" validate:"omitempty,min=3,max=100"`
}

type Output struct {
	Success bool `json:"success"`
}

package user

type Input struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Output struct {
	Success bool `json:"success"`
}

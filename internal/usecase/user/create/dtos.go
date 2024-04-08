package create

type Input struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Output struct {
	ID int64 `json:"id"`
}

package create

type Input struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type Output struct {
	ID int64 `json:"id"`
}

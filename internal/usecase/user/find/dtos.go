package find

import "time"

type Input struct {
	ID int64 `json:"id"`
}

type Output struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	LastLogin  time.Time `json:"last_login"`
}

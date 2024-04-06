package entity

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrNameRequired     = errors.New("name is required")
	ErrEmailRequired    = errors.New("email is required")
	ErrPasswordRequired = errors.New("password is required and cant't be blank")
	ErrPasswordLen      = errors.New("password must have at least caracters")
)

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
	LastLogin  time.Time `json:"last_login"`
}

func NewUser(name, email, password string) (*User, error) {
	user := User{
		Name:       strings.ToUpper(name),
		Email:      strings.ToLower(email),
		ModifiedAt: time.Now(),
	}

	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}

	err = user.Validate()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return ErrNameRequired
	}

	if u.Email == "" {
		return ErrEmailRequired
	}

	blankPassword := fmt.Sprintf("%x", (md5.Sum([]byte(""))))

	if u.Password == blankPassword {
		return ErrPasswordRequired
	}
	return nil
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	if len(password) < 6 {
		return ErrPasswordLen
	}

	u.Password = fmt.Sprintf("%x", (md5.Sum([]byte(password))))
	return nil
}

func (u *User) Update(name, email, password string) error {
	u.Name = strings.ToUpper(name)
	u.Email = strings.ToLower(email)
	err := u.SetPassword(password)
	if err != nil {
		return err
	}

	return u.Validate()
}

package entity

import (
	"errors"
	"github.com/andreluizmicro/go-driver-api/internal/domain/exception"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const PasswordMinLength = 8

var (
	ErrEmailRequired    = errors.New("email is required")
	ErrPasswordRequired = errors.New("password is required and can't be blank")
	ErrPasswordLen      = errors.New("password must have at least characters")
)

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
	LastLogin  time.Time `json:"last_login"`
}

func NewUser(name, email, password string) (*User, error) {

	user := User{
		Name:  strings.ToUpper(name),
		Email: strings.ToLower(email),
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
		return exception.ErrNameRequired
	}

	if u.Email == "" {
		return ErrEmailRequired
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(""), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	blankPassword := string(hash)

	if u.Password == blankPassword {
		return ErrPasswordRequired
	}
	return nil
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	if len(password) < PasswordMinLength {
		return ErrPasswordLen
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) Update(name, email string) error {
	u.Name = strings.ToUpper(name)
	u.Email = strings.ToLower(email)
	return u.Validate()
}

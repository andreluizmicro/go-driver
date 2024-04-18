package entity

import (
	"crypto/md5"
	"errors"
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type testcase struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}

	t.Run("test create user", func(t *testing.T) {
		testCases := []testcase{
			{Name: "", Email: "batman", Password: "123456789", ExpectedError: ErrNameRequired},
			{Name: "Batman da Silva", Email: "batman", Password: "12345678", ExpectedError: nil},
			{Name: "Super Man da Silva", Email: "batman", Password: "", ExpectedError: ErrPasswordRequired},
			{Name: "Hulk Da Silva", Email: "", Password: "123456789", ExpectedError: ErrEmailRequired},
			{Name: "Maria Da Silva", Email: "maria", Password: "", ExpectedError: ErrPasswordRequired},
			{Name: "José Da Silva", Email: "jose", Password: "123", ExpectedError: ErrPasswordLen},
		}

		for _, item := range testCases {
			user, err := NewUser(item.Name, item.Email, item.Password)
			if err != nil && !errors.Is(err, item.ExpectedError) {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}

			if user != nil {
				user.Password = fmt.Sprintf("%x", md5.Sum([]byte("")))
				if err = user.Validate(); err != nil && !errors.Is(err, ErrPasswordRequired) {
					t.Errorf("Expected %f but got %f", ErrPasswordRequired, err)
				}
			}
		}
	})
}

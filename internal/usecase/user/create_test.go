package user

import (
	"github.com/andreluizmicro/go-driver-api/internal/domain/exception"
	"github.com/andreluizmicro/go-driver-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateNewUser(t *testing.T) {
	repositoryMock := &mocks.UserRepositoryMock{}
	useCase := NewCreateUser(repositoryMock)

	t.Run("Should Return New User", func(t *testing.T) {
		var id int64 = 25
		repositoryMock.On("Create", mock.Anything).Return(&id, nil).Once()
		output, err := useCase.Execute(CreateInput{Name: "John", Email: "john@gmail.com", Password: "123456789"})
		assert.Nil(t, err)
		assert.Equal(t, id, output.ID)
		repositoryMock.AssertNumberOfCalls(t, "Create", 1)
	})

	t.Run("Should Return ErrUserAlreadyExists", func(t *testing.T) {
		var id int64 = 0
		repositoryMock.On("Create", mock.Anything).Return(&id, exception.ErrUserAlreadyExists).Once()
		output, err := useCase.Execute(CreateInput{Name: "John Doe", Email: "john@gmail.com", Password: "123456789"})
		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "user already exists")
		repositoryMock.AssertNumberOfCalls(t, "Create", 2)
	})

	t.Run("Should Return ErrNameRequired", func(t *testing.T) {
		output, err := useCase.Execute(CreateInput{Name: "", Email: "john@gmail.com", Password: "123456789"})
		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "name is required")
	})

	repositoryMock.AssertExpectations(t)
}

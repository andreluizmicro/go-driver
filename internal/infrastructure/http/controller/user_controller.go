package controller

import (
	"errors"
	deleteUseCase "github.com/andreluizmicro/go-driver-api/internal/usecase/user/delete"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/update"
	"net/http"

	"github.com/andreluizmicro/go-driver-api/internal/domain/exception"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/create"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/find"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUseCase *create.User
	findUseCase   *find.User
	updateUseCase *update.User
	deleteUseCase *deleteUseCase.User
}

func NewUserController(
	createUseCase *create.User,
	findUseCase *find.User,
	updateUseCase *update.User,
	deleteUseCase *deleteUseCase.User,
) *UserController {
	return &UserController{
		createUseCase: createUseCase,
		findUseCase:   findUseCase,
		updateUseCase: updateUseCase,
		deleteUseCase: deleteUseCase,
	}
}

func (us *UserController) Create(c *gin.Context) {
	var input create.Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": err.Error()})
		return
	}

	output, err := us.createUseCase.Execute(input)
	if err != nil {
		if errors.Is(err, exception.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (us *UserController) FindById(c *gin.Context) {
	var input find.Input

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	output, err := us.findUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &output)
}

func (us *UserController) Update(c *gin.Context) {
	var input update.Input

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	output, err := us.updateUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &output)
}

func (us *UserController) Delete(c *gin.Context) {
	var input deleteUseCase.Input

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err := us.deleteUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

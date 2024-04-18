package controller

import (
	"net/http"

	"github.com/andreluizmicro/go-driver-api/internal/domain/exception"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/create"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/find"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUseCase *create.CreateUser
	findUseCase   *find.FindUser
}

func NewUserController(
	createUseCase *create.CreateUser,
	findUseCase *find.FindUser,
) *UserController {
	return &UserController{
		createUseCase: createUseCase,
		findUseCase:   findUseCase,
	}
}

func (us *UserController) Create(c *gin.Context) {
	var input create.Input

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": err.Error()})
		return
	}

	output, err := us.createUseCase.Execute(input)
	if err != nil {
		if err == exception.ErrUserAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"msg": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (us *UserController) FindById(c *gin.Context) {
	var input find.Input

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": err.Error()})
		return
	}

	output, err := us.findUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &output)
}

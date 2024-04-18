package controller

import (
	"errors"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user"
	"net/http"

	"github.com/andreluizmicro/go-driver-api/internal/domain/exception"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUser *user.CreateUser
	findUser   *user.FindUser
	updateUser *user.UpdateUser
	deleteUser *user.DeleteUser
}

func NewUserController(
	createUser *user.CreateUser,
	findUser *user.FindUser,
	updateUser *user.UpdateUser,
	deleteUser *user.DeleteUser,
) *UserController {
	return &UserController{
		createUser: createUser,
		findUser:   findUser,
		updateUser: updateUser,
		deleteUser: deleteUser,
	}
}

func (us *UserController) Create(c *gin.Context) {
	var input user.CreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": err.Error()})
		return
	}

	output, err := us.createUser.Execute(input)
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
	var input user.FindInput

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	output, err := us.findUser.Execute(input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &output)
}

func (us *UserController) Update(c *gin.Context) {
	var input user.UpdateInput

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	output, err := us.updateUser.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &output)
}

func (us *UserController) Delete(c *gin.Context) {
	var input user.DeleteInput

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	output := us.deleteUser.Execute(input)
	if !output.Success {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when try delete user"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

package controller

import (
	"github.com/andreluizmicro/go-driver-api/internal/usecase/folder"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FolderController struct {
	createFolder *folder.CreateFolder
	updateFolder *folder.UpdateFolder
}

func NewFolderController(
	createFolder *folder.CreateFolder,
	updateFolder *folder.UpdateFolder,
) *FolderController {
	return &FolderController{
		createFolder: createFolder,
		updateFolder: updateFolder,
	}
}

func (fc *FolderController) Create(c *gin.Context) {
	var input folder.CreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	output, err := fc.createFolder.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (fc *FolderController) Update(c *gin.Context) {
	var input folder.UpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	_, err := fc.updateFolder.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, "")
}

package controller

import (
	"github.com/andreluizmicro/go-driver-api/internal/usecase/folder"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FolderController struct {
	createFolder *folder.CreateFolder
	updateFolder *folder.UpdateFolder
	deleteFolder *folder.DeleteFolder
}

func NewFolderController(
	createFolder *folder.CreateFolder,
	updateFolder *folder.UpdateFolder,
	deleteFolder *folder.DeleteFolder,
) *FolderController {
	return &FolderController{
		createFolder: createFolder,
		updateFolder: updateFolder,
		deleteFolder: deleteFolder,
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

func (fc *FolderController) Delete(c *gin.Context) {
	var input folder.DeleteInput
	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	output, err := fc.deleteFolder.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusNoContent, output)
}

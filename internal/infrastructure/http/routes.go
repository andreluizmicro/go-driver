package http

import (
	"fmt"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/http/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	userController *controller.UserController,
	folderController *controller.FolderController,
	port string,
) {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		users.GET("/", userController.FindAll)
		users.POST("/", userController.Create)
		users.GET("/:id", userController.FindById)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)

		folders := v1.Group("/folders")
		folders.POST("/", folderController.Create)
		folders.PUT("/:id", folderController.Update)
		folders.DELETE("/:id", folderController.Delete)
	}

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return
	}
}

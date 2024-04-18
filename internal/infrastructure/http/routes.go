package http

import (
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/http/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(userController *controller.UserController) {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		users.POST("/", userController.Create)
		users.GET("/:id", userController.FindById)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}

	err := router.Run(":9000")
	if err != nil {
		return
	}
}

package router

import (
	"blogs/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)

	userRouter := router.Group("/users")
	{
		userRouter.GET("/", controllers.GetUsers)
		userRouter.GET("/:id", controllers.GetUserByID)
		userRouter.POST("/", controllers.CreateUser)
		userRouter.PUT("/:id", controllers.UpdateUser)
		userRouter.DELETE("/:id", controllers.DeleteUser)
	}
}

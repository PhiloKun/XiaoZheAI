package router

import (
	"xiaozhe-ai-backend/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", handler.HealthCheck)
	}

	user := router.Group("/api/v1/user")
	{
		user.POST("/addUser", handler.AddUser)
		user.GET("/listUser", handler.ListUser)
		user.PUT("/updateUser", handler.UpdateUser)
		user.DELETE("/deleteUser", handler.DeleteUser)
	}
}

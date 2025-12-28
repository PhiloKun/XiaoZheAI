package router

import (
	"github.com/philokun/xiaozhe-ai-backend/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", handler.HealthCheck)
	}

	user := router.Group("/api/v1/user")
	{
		user.POST("/userRegister", handler.UserRegister)
		user.GET("/listUser", handler.ListUser)
		user.PUT("/updateUser", handler.UpdateUser)
		user.DELETE("/deleteUser", handler.DeleteUser)
	}

}

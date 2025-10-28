package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	InitRoutes(router)
	router.Run(":8080")
}

package router

import (
	"github.com/philokun/xiaozhe-ai-backend/config"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	InitRoutes(router)
	router.Run(config.GetServerAddress())
}

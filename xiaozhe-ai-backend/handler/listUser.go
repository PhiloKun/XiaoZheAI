package handler

import "github.com/gin-gonic/gin"

func ListUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ListUser",
	})
}

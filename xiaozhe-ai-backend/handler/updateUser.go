package handler

import "github.com/gin-gonic/gin"

func UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "UpdateUser",
	})
}

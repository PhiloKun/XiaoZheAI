package handler

import "github.com/gin-gonic/gin"

func UserRegister(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "AddUser",
	})
}

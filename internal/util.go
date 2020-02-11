package internal

import "github.com/gin-gonic/gin"

func NotFoundRoute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"staus": "NotFound",
	})
}

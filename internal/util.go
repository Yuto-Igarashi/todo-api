package internal

import "github.com/gin-gonic/gin"

//ページが存在しない場合にstaus:NotFoundを返す
func NotFoundRoute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"staus": "NotFound",
	})
}

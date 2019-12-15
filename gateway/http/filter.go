package http

import (
	"github.com/gin-gonic/gin"
)

func GlobalAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "token,content-type")
		if context.Request.Method == "OPTIONS" {
			context.Abort()
		}
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8090")	// 只允许http://localhost:8090 同源请求；区分localhost和127.0.0.1
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081") // 只允许http://localhost:8090 同源请求
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")                      //缓存时间
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")                    // 允许所有源的方法请求
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}

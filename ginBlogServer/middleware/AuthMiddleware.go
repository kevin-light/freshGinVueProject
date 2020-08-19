package middleware

import (
	"github.com/gin-gonic/gin"
	"lish.org/ginBlogs/common"
	"lish.org/ginBlogs/model"
	"net/http"
	"strings"
)

// gin的中间件= HandlerFunc返回fun
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		// !token.Valid 判断token无效
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}
		// 验证通过后获取claims 中的userID
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户不存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}
		// 用户存在 将user的信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}

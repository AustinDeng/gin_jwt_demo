package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gin_jwt_demo/utils"
)

// 校验请求头是否携带 `token` 字段, 并校验有效时间
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		claims, err := utils.ParseToken(token)

		if err != nil {
			if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

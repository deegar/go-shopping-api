package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sys_api/model"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		user := claims.(*model.CustomClaims)
		if user.AuthorityId != 2 {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "权限不足，禁止访问",
			})
			c.Abort()
		}
		c.Next()
	}
}

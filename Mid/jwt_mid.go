package mid

import (
	"ctf/logic"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token != "" {
			claims, _ := logic.Jwt_parse(token)
			if claims.Auth {
				logic.Jwt_update(token)
				c.Next()
				return
			}
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "no auth",
		})
		c.Abort()
	}
}

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "no token",
			})
			c.Abort()
			return
		}
		claims, err := logic.Jwt_parse(token)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "parse token failed",
			})
			c.Abort()
			return
		}
		if time.Now().After(claims.ExpiresAt.Time) {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "token expired",
			})
			c.Abort()
			return
		}
		logic.Jwt_update(token)
		c.Set("username", claims.Username)
		c.Set("auth", claims.Auth)
		c.Next()
	}
}

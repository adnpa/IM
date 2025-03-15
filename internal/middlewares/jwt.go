package middlewares

import (
	"net/http"

	"github.com/adnpa/IM/internal/utils"
	"github.com/gin-gonic/gin"
)

// var (
// 	TokenExpired     = errors.New("Token is expired")
// 	TokenNotValidYet = errors.New("Token not active yet")
// 	TokenMalformed   = errors.New("That's not even a token")
// 	TokenInvalid     = errors.New("Couldn't handle this token:")
// )

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "请登录"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "未登陆")
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Set("userId", claims.ID)

		c.Next()
	}
}

//https://oauth.net/2/

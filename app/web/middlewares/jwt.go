package middlewares

import (
	"net/http"
	"strings"

	"github.com/adnpa/IM/app/web/constant"
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
		authHeader := c.Request.Header.Get("authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
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

		c.Set(constant.CLAIMS_KEY, claims)
		c.Set(constant.USER_ID_KEY, claims.UID)

		c.Next()
	}
}

//https://oauth.net/2/

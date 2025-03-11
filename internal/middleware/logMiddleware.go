package middleware

import (
	"bytes"
	"io"

	"github.com/adnpa/IM/pkg/common/logger"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 读取请求体（以便后续打印）
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}

		// 将请求体重新写回，以便后续处理
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		logger.Infof("api argues", "query args", c.Request.URL.Path, "body args", string(bodyBytes))

		// 继续处理请求
		c.Next()
	}
}

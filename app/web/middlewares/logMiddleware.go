package middlewares

import (
	"bytes"
	"io"

	"github.com/adnpa/IM/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		logger.Info("Api Request", zap.String("query args", c.Request.URL.Path), zap.String("body args", string(bodyBytes)))

		// 继续处理请求
		c.Next()
	}
}

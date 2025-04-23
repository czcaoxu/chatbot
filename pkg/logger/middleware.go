package logger

import (
	"github.com/gin-gonic/gin"
	"time"
)

// GinLogger 是一个 Gin 中间件，用于记录每个请求的信息
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		Sugar().Infow("incoming request",
			"status", status,
			"method", c.Request.Method,
			"path", path,
			"ip", c.ClientIP(),
			"latency", latency.String(),
			"user_agent", c.Request.UserAgent(),
			"errors", c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}

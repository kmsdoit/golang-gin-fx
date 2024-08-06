package logger

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		c.Writer.Write([]byte("latency: " + latency.String()))
	}
}

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weirubo/learn_go/lesson64/prom"
)

func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		code := fmt.Sprintf("%d", c.Writer.Status())
		method := c.Request.Method
		labels := []string{"example-service", code, c.FullPath(), method}
		prom.RequestCounter.WithLabelValues(labels...).Inc()
	}
}

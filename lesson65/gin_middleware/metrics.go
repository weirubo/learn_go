package gin_middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/weirubo/learn_go/lesson65/prometheus_metrics"
)

func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		host := c.RemoteIP()
		code := fmt.Sprintf("%d", c.Writer.Status())
		method := c.Request.Method
		labelsByHttpReqs := []string{host, code, c.FullPath(), method}
		prometheus_metrics.HttpReqs.WithLabelValues(labelsByHttpReqs...).Inc()
		prometheus_metrics.OpsQueued.With(prometheus.Labels{"type": "delete", "user": "alice"}).Inc()
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/weirubo/learn_go/lesson65/gin_middleware"
	"github.com/weirubo/learn_go/lesson65/prometheus_metrics"
)

func init() {
	// 注册
	reg := prometheus.NewRegistry()
	reg.MustRegister(prometheus_metrics.HttpReqs, prometheus_metrics.OpsQueued, prometheus_metrics.Latencies, prometheus_metrics.Temps)
}

func main() {
	r := gin.New()
	r.Use(gin_middleware.Metrics())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}

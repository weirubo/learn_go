package prom

import "github.com/prometheus/client_golang/prometheus"

var (
	labelNames     = []string{"service", "code", "path", "method"}
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_count_total",
			Help: "Total number of HTTP requests made.",
		}, labelNames,
	)
)

func init() {
	prometheus.MustRegister(RequestCounter)
}

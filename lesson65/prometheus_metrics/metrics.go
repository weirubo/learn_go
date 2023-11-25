package prometheus_metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// 标签名
	labelNamesByHttpReqs = []string{"server", "code", "path", "method"}
	// HttpReqs 实例化 CounterVec
	HttpReqs *prometheus.CounterVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
	},
		labelNamesByHttpReqs,
	)

	labelNamesByOpsQueued = []string{
		// Which user has requested the operation?
		"user",
		// Of what type is the operation?
		"type",
	}
	OpsQueued = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			// fqName = strings.Join([]string{namespace, subsystem, name}, "_")
			//Namespace: "our_company",
			//Subsystem: "blob_storage",
			Name: "ops_queued",
			// fqName 相同的指标，Help 信息也需要相同。
			Help: "Number of blob storage operations waiting to be processed, partitioned by user and type.",
		},
		labelNamesByOpsQueued,
	)

	labelNamesByHistogram = []string{"method", "code"}
	Latencies             = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Tracks the latencies for HTTP requests.",
			Buckets: []float64{0.99, 0.9, 0.5},
		},
		labelNamesByHistogram,
	)

	labelNamesByTemps = []string{"species"}
	Temps             = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "pond_temperature_celsius",
			Help:       "The temperature of the frog pond.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		labelNamesByTemps,
	)
)

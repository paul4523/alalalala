package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		
		Help: "Duration of HTTP requests",
	}, []string{"endpoint"})

	RequestCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "Number of HTTP requests",
	}, []string{"endpoint"})

	ApiDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "api_request_duration_seconds",
		Help: "Duration of API requests",
	})

	CacheDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "cache_request_duration_seconds",
		Help: "Duration of cache requests",
	})
)

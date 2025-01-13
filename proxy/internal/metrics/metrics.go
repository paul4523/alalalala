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

	CacheDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "cache_request_duration_seconds",
		Help: "Duration of cache requests",
	}, []string{"method"})

	DBDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "db_request_duration_seconds",
		Help: "Duration of database requests",
	}, []string{"method"})

	APIDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "api_request_duration_seconds",
		Help: "Duration of external API requests",
	}, []string{"method"})
)

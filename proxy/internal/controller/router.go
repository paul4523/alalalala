package controller

import (
	"net/http"
	"time"

	"proxy/internal/metrics"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Handle("/metrics", promhttp.Handler())

	r.Post("/api/address/search", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.AddressSearch(w, r)
		metrics.RequestDuration.WithLabelValues("/api/address/search").Observe(time.Since(start).Seconds())
		metrics.RequestCount.WithLabelValues("/api/address/search").Inc()
	})

	return r
}

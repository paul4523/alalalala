package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/api/address/search", func(w http.ResponseWriter, r *http.Request) {
		h.AddressSearch(w, r)
	})

	r.Handle("/metrics", promhttp.Handler())

	return r
}

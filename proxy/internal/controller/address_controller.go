package controller

import (
	"encoding/json"
	"net/http"

	"proxy/internal/model"
)

type Responder interface {
	AddressSearch(input string) ([]*model.Address, error)
	Cash(input string) ([]*model.Address, error)
}
type Handler struct {
	Responder Responder
}

func (h *Handler) AddressSearch(w http.ResponseWriter, r *http.Request) {
	var req model.RequestAddressSearch
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addresses, err := h.Responder.AddressSearch(req.Query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := model.ResponseAddress{Addresses: addresses}
	if err := response.Respond(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

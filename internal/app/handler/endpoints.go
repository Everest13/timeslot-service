package handler

import "net/http"

func (h *Handler) ConfigureHTTPEndpoints() map[string]map[string]http.HandlerFunc {
	return map[string]map[string]http.HandlerFunc{
		http.MethodPost: {
			"/slots": h.BookSlots,
		},
		http.MethodDelete: {
			"/slots": h.CancelSlots,
		},
	}
}

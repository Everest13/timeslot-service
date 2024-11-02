package handler

import (
	"encoding/json"
	"net/http"
	"timeslot-service/internal/models"
)

// слой контроллера
func (h *Handler) BookSlots(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slots, err := h.timeSlotService.BookSlots(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(slots) == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

package handler

import (
	"timeslot-service/internal/service/timeslot"
)

type Handler struct {
	timeSlotService *timeslot.Service
}

func NewHandler(timeSlotService *timeslot.Service) *Handler {
	return &Handler{timeSlotService: timeSlotService}
}

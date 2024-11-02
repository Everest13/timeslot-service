package repository

import (
	"time"
	"timeslot-service/internal/models"
)

type DB interface {
	GetRequiredSlots(orderFrom, orderTo time.Time) ([]models.TimeSlot, error)
	UpdateTimeSlot(slots []models.TimeSlot) error
	SetReservations(reservations []models.Reservation) error
	Lock()
	Unlock()
}

type Repository struct {
	db DB
}

func NewRepository(db DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) TakeTimeSlot(order models.Order) ([]models.TimeSlot, error) {
	r.db.Lock() //имитация транзакции
	defer r.db.Unlock()

	//1 - timeslots select
	requiredSlots, err := r.db.GetRequiredSlots(order.From, order.To)
	if err != nil {
		return nil, err
	}

	if len(requiredSlots) == 0 {
		return nil, nil
	}

	for i := 0; i < len(requiredSlots); i++ {
		if requiredSlots[i].Capacity < order.Capacity {
			return nil, nil
		}
	}

	firstSlot := requiredSlots[0]
	lastSlot := requiredSlots[len(requiredSlots)-1]
	if !firstSlot.From.Before(order.From) || !lastSlot.To.After(order.To) {
		return nil, nil
	}

	//2 - timeslots update
	err = r.db.UpdateTimeSlot(requiredSlots)
	if err != nil {
		return nil, err
	}

	//3 - reservation insert батч записей
	reservations := make([]models.Reservation, 0, len(requiredSlots))
	for _, slot := range requiredSlots {
		reservations = append(reservations, models.Reservation{
			TimeslotID: slot.ID,
			RequestID:  order.RequestID,
			Capacity:   order.Capacity,
		})
	}

	err = r.db.SetReservations(reservations)
	if err != nil {
		return nil, err
	}

	return requiredSlots, nil
}

// потенциальное расширение для получения данных аналитиками
func (r *Repository) GetReservation(order models.Order) ([]models.Reservation, error) {
	return nil, nil
}

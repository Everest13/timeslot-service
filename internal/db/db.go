package db

import (
	"sync"
	"time"
	"timeslot-service/internal/models"
)

type DB struct {
	Mutex             sync.Mutex
	reservationsTable []models.Reservation
	timeSlotsTable    map[string]models.TimeSlot
}

func NewDB() *DB {
	return &DB{
		timeSlotsTable:    map[string]models.TimeSlot{},
		reservationsTable: []models.Reservation{},
	}
}

func (db *DB) Lock() {
	db.Mutex.Lock()
}

func (db *DB) Unlock() {
	db.Mutex.Unlock()
}

func (db *DB) GetRequiredSlots(orderFrom, orderTo time.Time) ([]models.TimeSlot, error) {
	slots := db.timeSlotsTable
	requiredSlots := []models.TimeSlot{}
	for _, slot := range slots {
		if (slot.From.After(orderFrom) && slot.From.Before(orderTo)) ||
			(slot.To.After(orderFrom) && slot.To.Before(orderTo)) {
			requiredSlots = append(requiredSlots, slot) // часть или весь слот попадает внутрь order'а
		} else if slot.From.Before(orderFrom) && slot.To.After(orderTo) {
			requiredSlots = append(requiredSlots, slot) // весь order попадает в один слот
		}
	}

	return requiredSlots, nil
}

func (db *DB) UpdateTimeSlot(slots []models.TimeSlot) error {
	for _, slot := range slots {
		db.timeSlotsTable[slot.ID] = slot
	}

	return nil
}

func (db *DB) SetReservations(reservations []models.Reservation) error {
	db.reservationsTable = append(db.reservationsTable, reservations...)

	return nil
}

package repository

import (
	"testing"
	"time"
	"timeslot-service/internal/models"
)

func TestTakeTimeSlot(t *testing.T) {
	tests := []struct {
		name                   string
		order                  models.Order
		getRequiredSlotsResult []models.TimeSlot
		updateTimeSlotResult   error
		setReservations        error
	}{
		{
			name: "success",
			order: models.Order{
				RequestID: "request1",
				From:      time.Date(2024, 10, 01, 11, 30, 00, 0, time.UTC),
				To:        time.Date(2024, 10, 01, 13, 30, 00, 0, time.UTC),
				Capacity:  1,
			},
			getRequiredSlotsResult: []models.TimeSlot{
				{
					ID:       "slot1",
					From:     time.Date(2024, 10, 01, 11, 00, 00, 0, time.UTC),
					To:       time.Date(2024, 10, 01, 11, 59, 00, 0, time.UTC),
					Capacity: 1,
				},
				{
					ID:       "slot2",
					From:     time.Date(2024, 10, 01, 12, 00, 00, 0, time.UTC),
					To:       time.Date(2024, 10, 01, 12, 59, 00, 0, time.UTC),
					Capacity: 1,
				},
				{
					ID:       "slot3",
					From:     time.Date(2024, 10, 01, 13, 00, 00, 0, time.UTC),
					To:       time.Date(2024, 10, 01, 13, 59, 00, 0, time.UTC),
					Capacity: 1,
				},
				{
					ID:       "slot4",
					From:     time.Date(2024, 10, 01, 14, 00, 00, 0, time.UTC),
					To:       time.Date(2024, 10, 01, 14, 59, 00, 0, time.UTC),
					Capacity: 1,
				},
			},
			updateTimeSlotResult: nil,
			setReservations:      nil,
		},
		{
			name: "not enough slots",
			order: models.Order{
				RequestID: "request2",
				From:      time.Date(2024, 10, 01, 11, 30, 00, 0, time.UTC),
				To:        time.Date(2024, 10, 01, 13, 30, 00, 0, time.UTC),
				Capacity:  2,
			},
			getRequiredSlotsResult: []models.TimeSlot{
				{
					ID:       "slot1",
					From:     time.Date(2024, 10, 01, 11, 00, 00, 0, time.UTC),
					To:       time.Date(2024, 10, 01, 11, 59, 00, 0, time.UTC),
					Capacity: 2,
				},
				{
					ID:       "slot2",
					From:     time.Date(2024, 10, 01, 12, 00, 00, 0, time.UTC),
					To:       time.Date(2024, 10, 01, 12, 59, 00, 0, time.UTC),
					Capacity: 2,
				},
			},
			updateTimeSlotResult: nil,
			setReservations:      nil,
		},
		{
			name: "not enough slot's capacity",
			order: models.Order{
				RequestID: "request3",
				From:      time.Date(2024, 10, 01, 11, 30, 00, 0, time.UTC),
				To:        time.Date(2024, 10, 01, 13, 30, 00, 0, time.UTC),
				Capacity:  2,
			},
			getRequiredSlotsResult: []models.TimeSlot{
				{
					ID:       "slot1",
					From:     time.Date(2024, 10, 01, 11, 00, 00, 0, time.UTC),
					To:       time.Date(2024, 10, 01, 11, 59, 00, 0, time.UTC),
					Capacity: 2,
				},
				{
					ID:       "slot2",
					From:     time.Date(2024, 10, 01, 12, 00, 00, 0, time.UTC),
					To:       time.Date(2024, 10, 01, 12, 59, 00, 0, time.UTC),
					Capacity: 1,
				},
				{
					ID:       "slot3",
					From:     time.Date(2024, 10, 01, 13, 00, 00, 0, time.UTC),
					To:       time.Date(2024, 10, 01, 13, 59, 00, 0, time.UTC),
					Capacity: 2,
				},
			},
			updateTimeSlotResult: nil,
			setReservations:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fakeDB := fakeDataBase{
				getRequiredSlotsResult: test.getRequiredSlotsResult,
				updateTimeSlotResult:   test.updateTimeSlotResult,
				setReservations:        test.setReservations,
			}

			repo := NewRepository(&fakeDB)
			slots, err := repo.TakeTimeSlot(test.order)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if len(slots) == 0 {
				t.Errorf("imposible set order: slots is not enough")
			}
		})
	}
}

type fakeDataBase struct {
	lockCallNum            int
	unLockCallNum          int
	getRequiredSlotsResult []models.TimeSlot
	updateTimeSlotResult   error
	setReservations        error
}

func (fdb *fakeDataBase) GetRequiredSlots(orderFrom, orderTo time.Time) ([]models.TimeSlot, error) {
	return fdb.getRequiredSlotsResult, nil
}

func (fdb *fakeDataBase) UpdateTimeSlot(slots []models.TimeSlot) error {
	return fdb.updateTimeSlotResult
}

func (fdb *fakeDataBase) SetReservations(reservations []models.Reservation) error {
	return fdb.setReservations
}

func (fdb *fakeDataBase) Lock() {}

func (fdb *fakeDataBase) Unlock() {}

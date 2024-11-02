package timeslot

import (
	"timeslot-service/internal/models"
	"timeslot-service/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) BookSlots(order models.Order) ([]models.TimeSlot, error) {
	slots, err := s.repo.TakeTimeSlot(order)
	if err != nil {
		return nil, err
	}

	return slots, nil
}

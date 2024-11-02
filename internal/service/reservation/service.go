package reservation

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

// пример получения данных для аналитиков
func (s *Service) GetReservation(order models.Order) ([]models.Reservation, error) {
	return s.repo.GetReservation(order)
}

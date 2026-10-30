package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	Triprepo domain.TripRepository
}

func NewTripService(Triprepo domain.TripRepository) *Service {
	return &Service{Triprepo: Triprepo}
}

func (s *Service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserID:   "",
		Status:   "pandding",
		RiddFare: fare,
	}

	return s.Triprepo.CreateTrip(ctx, t)
}

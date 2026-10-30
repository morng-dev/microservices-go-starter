package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string
	PackageSlut       string //EX: van, luxury, sedan
	TotalPriceUnCents float64
	ExpriesAt         time.Time
}

type TripModel struct {
	ID       primitive.ObjectID
	UserID   string
	Status   string
	RiddFare *RideFareModel
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
}

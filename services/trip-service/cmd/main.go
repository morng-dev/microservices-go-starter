package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInmemRepository()
	service := service.NewTripService(inmemRepo)
	fare := &domain.RideFareModel{
		UserID: "42",
	}

	trip, err := service.CreateTrip(ctx, fare)
	if err != nil {
		log.Printf("error :%v", err)
	}
	log.Println(trip)
	for {
		time.Sleep(time.Second)
	}
}

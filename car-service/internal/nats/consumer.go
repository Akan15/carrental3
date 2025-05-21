package nats

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Akan15/carrental3/car-service/internal/usecase"

	"github.com/nats-io/nats.go"
)

type RentalCreatedEvent struct {
	RentalID string `json:"rental_id"`
	CarID    string `json:"car_id"`
	Status   string `json:"status"`
}

func SubscribeToRentalCreated(natsURL string, carUC usecase.CarUsecase) error {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return err
	}

	_, err = nc.Subscribe("rental.created", func(msg *nats.Msg) {
		var event RentalCreatedEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("Invalid message:", err)
			return
		}
		err := carUC.ChangeStatus(context.Background(), event.CarID, event.Status)
		if err != nil {
			log.Println("Failed to update car status:", err)
		}
	})
	return err
}

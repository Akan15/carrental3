package nats

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type RentalCreatedEvent struct {
	RentalID string `json:"rental_id"`
	CarID    string `json:"car_id"`
	Status   string `json:"status"` // например, "occupied"
}

type Publisher struct {
	conn *nats.Conn
}

func NewPublisher(natsURL string) (*Publisher, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	return &Publisher{conn: nc}, nil
}

func (p *Publisher) PublishRentalCreated(event RentalCreatedEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return p.conn.Publish("rental.created", data)
}

func (p *Publisher) Close() {
	p.conn.Close()
}

package main

import (
	"log"

	"github.com/Akan15/carrental3/consumer-service/internal/subscriber"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://nats:4222") // NATS hostname из docker-compose
	if err != nil {
		log.Fatalf("❌ Failed to connect to NATS: %v", err)
	}
	defer nc.Drain()

	log.Println("✅ Connected to NATS")

	subscriber.SubscribeToUserCreated(nc)

	select {} // Бесконечное ожидание
}

package nats

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

var conn *nats.Conn

func InitPublisher() {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = "nats://nats:4222" // имя контейнера в docker-compose
	}
	var err error
	conn, err = nats.Connect(url)
	if err != nil {
		log.Fatalf("🔴 Failed to connect to NATS: %v", err)
	}
	log.Println("✅ Connected to NATS")
}

func PublishUserCreated(email string) {
	if conn == nil {
		log.Println("⚠️ NATS not connected")
		return
	}
	err := conn.Publish("user.created", []byte(email))
	if err != nil {
		log.Printf("🔴 Failed to publish: %v", err)
	}
}

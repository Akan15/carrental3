package nats

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

func InitPublisher() {
	var err error
	nc, err = nats.Connect(os.Getenv("NATS_URL")) // например: nats://nats:4222
	if err != nil {
		log.Fatalf("❌ Failed to connect to NATS: %v", err)
	}
	log.Println("✅ Connected to NATS")
}

func PublishUserCreated(email string) {
	if nc == nil {
		log.Println("❌ NATS not initialized")
		return
	}
	err := nc.Publish("user.created", []byte(email))
	if err != nil {
		log.Printf("❌ Failed to publish user.created: %v", err)
	} else {
		log.Printf("📨 Published user.created for %s", email)
	}
}

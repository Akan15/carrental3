package subscriber

import (
	"log"

	"github.com/nats-io/nats.go"
)

func SubscribeToUserCreated(nc *nats.Conn) {
	_, err := nc.Subscribe("user.created", func(msg *nats.Msg) {
		log.Printf("📨 Received user.created event: %s", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("❌ Failed to subscribe: %v", err)
	}

	log.Println("✅ Subscribed to 'user.created'")
}

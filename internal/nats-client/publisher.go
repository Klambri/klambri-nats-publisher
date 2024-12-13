package natsclient

import (
	"log"

	"github.com/nats-io/nats.go"
)

func publish(nc *nats.Conn, subject string, data []byte) error {
	err := nc.Publish(subject, data)
	if err != nil {
		log.Println("Failed to publish message:", err)
		return err
	}
	return nil
}

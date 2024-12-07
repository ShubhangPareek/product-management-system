package services

import (
	"testing"

	"github.com/rabbitmq/amqp091-go"
)

func TestRabbitMQConnection(t *testing.T) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
}

package test

import (
	"testing"

	"github.com/rabbitmq/amqp091-go"
)

func TestRabbitMQConsumer(t *testing.T) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"image_processing", false, false, false, false, nil,
	)
	if err != nil {
		t.Fatalf("Failed to declare a queue: %v", err)
	}

	t.Log("RabbitMQ connection and queue declaration successful!")
}

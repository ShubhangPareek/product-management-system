package services

import (
	"testing"

	"github.com/rabbitmq/amqp091-go"
)

func TestRabbitMQProducer(t *testing.T) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("Failed to open a RabbitMQ channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("image_processing", false, false, false, false, nil)
	if err != nil {
		t.Fatalf("Failed to declare a queue: %v", err)
	}

	err = ch.Publish("", q.Name, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte("test message"),
	})
	if err != nil {
		t.Fatalf("Failed to publish a message: %v", err)
	}

	t.Log("RabbitMQ producer test passed")
}

func TestRabbitMQConsumer(t *testing.T) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("Failed to open a RabbitMQ channel: %v", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume("image_processing", "", true, false, false, false, nil)
	if err != nil {
		t.Fatalf("Failed to consume messages: %v", err)
	}

	for msg := range msgs {
		t.Logf("Received message: %s", msg.Body)
		break
	}
}

package services

import (
	"log"
	"strings"

	"github.com/rabbitmq/amqp091-go"
)

func StartImageProcessor() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"image_processing", // Queue name
		false,              // Durable
		false,              // Delete when unused
		false,              // Exclusive
		false,              // No-wait
		nil,                // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer tag
		true,   // Auto-acknowledge
		false,  // Exclusive
		false,  // No local
		false,  // No wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Printf("Waiting for messages on queue: %s", q.Name)

	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
		imageURLs := strings.Split(string(msg.Body), ",")
		for _, imageURL := range imageURLs {
			log.Printf("Processing image: %s", imageURL)
			// Add image processing logic here
		}
	}
}

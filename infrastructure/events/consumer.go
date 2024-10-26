package events

import (
	"encoding/json"
	"log"
	"ticketing-system/domain/entities"
	"ticketing-system/infrastructure/repository"

	"github.com/streadway/amqp"
)

type EventConsumer struct {
	esRepo  repository.TicketQueryRepository
	channel *amqp.Channel
}

// NewEventConsumer creates a new EventConsumer instance
func NewEventConsumer(esRepo repository.TicketQueryRepository, channel *amqp.Channel) *EventConsumer {
	return &EventConsumer{esRepo: esRepo, channel: channel}
}

func (c *EventConsumer) StartConsumer() {
	log.Println("Starting RabbitMQ consumer for ticket.created events")

	// Declare the queue from which we will consume events
	queueName := "ticket.created"

	// Ensure the queue exists
	_, err := c.channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue %s: %v", queueName, err)
	}

	msgs, err := c.channel.Consume(
		queueName, // queue
		"",        // consumer tag
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to start consuming events: %v", err)
	}

	// Create a goroutine to handle messages
	go func() {
		for d := range msgs {
			log.Println("Received ticket.created event")

			var ticketEvent entities.Ticket
			err := json.Unmarshal(d.Body, &ticketEvent)
			if err != nil {
				log.Printf("Failed to unmarshal event: %v", err)
				continue
			}

			log.Printf("Processing event for ticket ID: %d", ticketEvent.ID)

			// Index the ticket in Elasticsearch
			err = c.esRepo.SaveToElasticsearch(&ticketEvent)
			if err != nil {
				log.Printf("Failed to save ticket to Elasticsearch: %v", err)
			} else {
				log.Printf("Successfully indexed ticket in Elasticsearch: %+v", ticketEvent)
			}
		}
	}()
}

// StopConsumer gracefully stops consuming messages from RabbitMQ
func (c *EventConsumer) StopConsumer() error {
	return c.channel.Cancel("", false)
}

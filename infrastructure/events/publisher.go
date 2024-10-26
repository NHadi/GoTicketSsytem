package events

import (
	"encoding/json"
	"log"
	"ticketing-system/domain/entities"

	"github.com/streadway/amqp"
)

// EventPublisher is an interface that defines methods for publishing events
type EventPublisher interface {
	PublishTicketCreated(t *entities.Ticket) error
	PublishTicketStatusUpdated(t *entities.Ticket) error
}

// RabbitMQEventPublisher implements EventPublisher for RabbitMQ
type RabbitMQEventPublisher struct {
	channel *amqp.Channel
}

// NewEventPublisher creates a new instance of RabbitMQEventPublisher
func NewEventPublisher(channel *amqp.Channel) *RabbitMQEventPublisher {
	return &RabbitMQEventPublisher{channel: channel}
}

// PublishTicketCreated publishes a "ticket.created" event to RabbitMQ
func (p *RabbitMQEventPublisher) PublishTicketCreated(t *entities.Ticket) error {
	queueName := "ticket.created"

	// Declare the queue to ensure it exists
	_, err := p.channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("Failed to declare queue %s: %v", queueName, err)
		return err
	}

	// Rest of the publish code remains unchanged
	body, err := json.Marshal(t)
	if err != nil {
		log.Printf("Failed to marshal ticket: %v", err)
		return err
	}

	err = p.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish ticket.created event: %v", err)
		return err
	}

	log.Printf("Published ticket.created event successfully for ticket ID: %d", t.ID)
	return nil
}

// PublishTicketStatusUpdated publishes a "ticket.status.updated" event to RabbitMQ
func (p *RabbitMQEventPublisher) PublishTicketStatusUpdated(t *entities.Ticket) error {
	log.Printf("Publishing ticket.status.updated event for ticket ID: %d", t.ID)

	body, err := json.Marshal(t)
	if err != nil {
		log.Printf("Failed to marshal ticket: %v", err)
		return err
	}

	// Publish the event to RabbitMQ
	err = p.channel.Publish(
		"",                      // exchange
		"ticket.status.updated", // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish ticket.status.updated event: %v", err)
		return err
	}

	log.Printf("Published ticket.status.updated event successfully for ticket ID: %d", t.ID)
	return nil
}

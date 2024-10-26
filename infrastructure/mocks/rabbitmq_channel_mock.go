package mocks

import (
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/mock"
)

// RabbitMQChannelMock is a mock implementation of the RabbitMQ channel interface for testing purposes.
type RabbitMQChannelMock struct {
	mock.Mock
}

// QueueDeclare mocks the QueueDeclare method in the RabbitMQ channel.
func (m *RabbitMQChannelMock) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	mockArgs := m.Called(name, durable, autoDelete, exclusive, noWait, args)
	return mockArgs.Get(0).(amqp.Queue), mockArgs.Error(1)
}

// Publish mocks the Publish method in the RabbitMQ channel.
func (m *RabbitMQChannelMock) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	mockArgs := m.Called(exchange, key, mandatory, immediate, msg)
	return mockArgs.Error(0)
}

package mocks

import (
	"ticketing-system/domain/entities"

	"github.com/stretchr/testify/mock"
)

// EventPublisherMock is a mock implementation of the EventPublisher interface.
type EventPublisherMock struct {
	mock.Mock
}

// PublishTicketCreated mocks the PublishTicketCreated method in EventPublisher.
func (m *EventPublisherMock) PublishTicketCreated(ticket *entities.Ticket) error {
	args := m.Called(ticket)
	return args.Error(0)
}

// PublishTicketStatusUpdated mocks the PublishTicketStatusUpdated method in EventPublisher.
func (m *EventPublisherMock) PublishTicketStatusUpdated(ticket *entities.Ticket) error {
	args := m.Called(ticket)
	return args.Error(0)
}

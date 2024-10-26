package mocks

import (
	"ticketing-system/domain/entities"

	"github.com/stretchr/testify/mock"
)

// TicketRepositoryMock is a mock implementation of the TicketRepository interface.
type TicketRepositoryMock struct {
	mock.Mock
}

// Save mocks the Save method in the TicketRepository.
func (m *TicketRepositoryMock) Save(ticket *entities.Ticket) error {
	args := m.Called(ticket)
	return args.Error(0)
}

// FindByID mocks the FindByID method in the TicketRepository.
func (m *TicketRepositoryMock) FindByID(id uint) (*entities.Ticket, error) {
	args := m.Called(id)
	if ticket, ok := args.Get(0).(*entities.Ticket); ok {
		return ticket, args.Error(1)
	}
	return nil, args.Error(1)
}

// UpdateStatus mocks the UpdateStatus method in the TicketRepository.
func (m *TicketRepositoryMock) UpdateStatus(ticket *entities.Ticket) error {
	args := m.Called(ticket)
	return args.Error(0)
}

package mocks

import (
	"ticketing-system/application/dto"
	"ticketing-system/domain/entities"

	"github.com/stretchr/testify/mock"
)

// TicketServiceMock is a mock implementation of the TicketService interface.
type TicketServiceMock struct {
	mock.Mock
}

// CreateTicket mocks the CreateTicket method.
func (m *TicketServiceMock) CreateTicket(title, message string, userID uint) (*entities.Ticket, error) {
	args := m.Called(title, message, userID)
	return args.Get(0).(*entities.Ticket), args.Error(1)
}

// UpdateTicketStatus mocks the UpdateTicketStatus method.
func (m *TicketServiceMock) UpdateTicketStatus(id string, newStatus entities.TicketStatus) error {
	args := m.Called(id, newStatus)
	return args.Error(0)
}

// GetTickets mocks the GetTickets method.
func (m *TicketServiceMock) GetTickets(filters *dto.FilterOptions, sort dto.SortOptions, pageSize, page int) ([]entities.Ticket, int, error) {
	args := m.Called(filters, sort, pageSize, page)
	return args.Get(0).([]entities.Ticket), args.Int(1), args.Error(2)
}

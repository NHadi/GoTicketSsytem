package mocks

import (
	"ticketing-system/application/dto"
	"ticketing-system/domain/entities"

	"github.com/stretchr/testify/mock"
)

// TicketQueryRepositoryMock is a mock implementation of the TicketQueryRepository interface.
type TicketQueryRepositoryMock struct {
	mock.Mock
}

// SaveToElasticsearch mocks the SaveToElasticsearch method.
func (m *TicketQueryRepositoryMock) SaveToElasticsearch(ticket *entities.Ticket) error {
	args := m.Called(ticket)
	return args.Error(0)
}

// Search mocks the Search method.
func (m *TicketQueryRepositoryMock) Search(query map[string]interface{}) ([]entities.Ticket, error) {
	args := m.Called(query)
	return args.Get(0).([]entities.Ticket), args.Error(1)
}

// GetTickets mocks the GetTickets method with filters, sorting, and pagination.
func (m *TicketQueryRepositoryMock) GetTickets(filters *dto.FilterOptions, sort dto.SortOptions, pageSize, page int) ([]entities.Ticket, int, error) {
	args := m.Called(filters, sort, pageSize, page)
	return args.Get(0).([]entities.Ticket), args.Int(1), args.Error(2)
}

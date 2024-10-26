package services_test

import (
	"testing"
	"ticketing-system/application/dto"
	"ticketing-system/application/services"
	"ticketing-system/domain/entities"
	"ticketing-system/infrastructure/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTicket(t *testing.T) {
	mockRepo := new(mocks.TicketRepositoryMock)
	mockPublisher := new(mocks.EventPublisherMock)
	mockQueryRepo := new(mocks.TicketQueryRepositoryMock)
	service := services.NewTicketService(mockRepo, mockPublisher, mockQueryRepo)

	title, message := "Sample Ticket", "This is a sample ticket message."
	userID := uint(1)

	mockRepo.On("Save", mock.AnythingOfType("*entities.Ticket")).Return(nil)
	mockPublisher.On("PublishTicketCreated", mock.AnythingOfType("*entities.Ticket")).Return(nil)

	ticket, err := service.CreateTicket(title, message, userID)

	assert.NoError(t, err)
	assert.Equal(t, title, ticket.Title)
	assert.Equal(t, entities.Open, ticket.Status)
}

func TestGetTickets(t *testing.T) {
	mockRepo := new(mocks.TicketRepositoryMock)
	mockPublisher := new(mocks.EventPublisherMock)
	mockQueryRepo := new(mocks.TicketQueryRepositoryMock)
	service := services.NewTicketService(mockRepo, mockPublisher, mockQueryRepo)

	filter := &dto.FilterOptions{FilterName: "CreatedAt", FilterType: "before", FilterValue: "2024-10-26T13:50:00Z"}
	sort := dto.SortOptions{SortName: "CreatedAt", SortDir: "asc"}

	mockRepo.On("GetTickets", filter, sort, 10, 1).Return([]entities.Ticket{
		{ID: 1, Title: "Ticket 1"},
		{ID: 2, Title: "Ticket 2"},
	}, 2, nil)

	tickets, totalCount, err := service.GetTickets(filter, sort, 10, 1)

	assert.NoError(t, err)
	assert.Equal(t, 2, totalCount)
	assert.Equal(t, "Ticket 1", tickets[0].Title)
}

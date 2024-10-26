package repository_test

import (
	"testing"
	"ticketing-system/domain/entities"
	"ticketing-system/infrastructure/mocks"
	"ticketing-system/infrastructure/repository"

	"github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveToElasticsearch(t *testing.T) {
	mockClient := new(mocks.ElasticsearchClientMock)
	mockIndexService := new(mocks.MockIndexService)

	repo := repository.NewElasticsearchTicketRepository(mockClient)

	ticket := &entities.Ticket{ID: 1, Title: "Test Ticket"}
	docID := "1"

	mockClient.On("Index").Return(mockIndexService)
	mockIndexService.On("Index", "tickets").Return(mockIndexService)
	mockIndexService.On("Id", docID).Return(mockIndexService)
	mockIndexService.On("BodyJson", ticket).Return(mockIndexService)
	mockIndexService.On("Do", mock.Anything).Return(&elastic.IndexResponse{}, nil)

	err := repo.SaveToElasticsearch(ticket)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
	mockIndexService.AssertExpectations(t)
}

func TestSearch(t *testing.T) {
	mockClient := new(mocks.ElasticsearchClientMock)
	mockSearchService := new(mocks.MockSearchService)

	repo := repository.NewElasticsearchTicketRepository(mockClient)

	query := map[string]interface{}{"Status": "Open"}

	mockClient.On("Search").Return(mockSearchService)
	mockSearchService.On("Index", "tickets").Return(mockSearchService)
	mockSearchService.On("Query", mock.AnythingOfType("*elastic.BoolQuery")).Return(mockSearchService)
	mockSearchService.On("Do", mock.Anything).Return(&elastic.SearchResult{}, nil)

	results, err := repo.Search(query)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	mockClient.AssertExpectations(t)
	mockSearchService.AssertExpectations(t)
}

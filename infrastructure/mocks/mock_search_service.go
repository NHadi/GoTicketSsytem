package mocks

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/mock"
)

// MockSearchService is a mock for the elastic.SearchService
type MockSearchService struct {
	mock.Mock
}

func (m *MockSearchService) Index(indices ...string) *elastic.SearchService {
	m.Called(indices)
	return &elastic.SearchService{}
}

func (m *MockSearchService) Query(query elastic.Query) *elastic.SearchService {
	m.Called(query)
	return &elastic.SearchService{}
}

func (m *MockSearchService) Sort(field string, ascending bool) *elastic.SearchService {
	m.Called(field, ascending)
	return &elastic.SearchService{}
}

func (m *MockSearchService) From(from int) *elastic.SearchService {
	m.Called(from)
	return &elastic.SearchService{}
}

func (m *MockSearchService) Size(size int) *elastic.SearchService {
	m.Called(size)
	return &elastic.SearchService{}
}

func (m *MockSearchService) Do(ctx context.Context) (*elastic.SearchResult, error) {
	args := m.Called(ctx)
	return args.Get(0).(*elastic.SearchResult), args.Error(1)
}

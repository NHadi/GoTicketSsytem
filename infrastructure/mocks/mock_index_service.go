package mocks

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/mock"
)

// MockIndexService is a mock for the elastic.IndexService
type MockIndexService struct {
	mock.Mock
}

func (m *MockIndexService) Index(index string) *elastic.IndexService {
	m.Called(index)
	return &elastic.IndexService{}
}

func (m *MockIndexService) Id(id string) *elastic.IndexService {
	m.Called(id)
	return &elastic.IndexService{}
}

func (m *MockIndexService) BodyJson(body interface{}) *elastic.IndexService {
	m.Called(body)
	return &elastic.IndexService{}
}

func (m *MockIndexService) Do(ctx context.Context) (*elastic.IndexResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).(*elastic.IndexResponse), args.Error(1)
}

package mocks

import (
	"github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/mock"
)

// ElasticsearchClientMock is a mock implementation of the ElasticsearchClient interface.
type ElasticsearchClientMock struct {
	mock.Mock
}

// Index mocks the Index method for Elasticsearch.
func (m *ElasticsearchClientMock) Index() *elastic.IndexService {
	args := m.Called()
	return args.Get(0).(*elastic.IndexService)
}

// Search mocks the Search method for Elasticsearch.
func (m *ElasticsearchClientMock) Search() *elastic.SearchService {
	args := m.Called()
	return args.Get(0).(*elastic.SearchService)
}

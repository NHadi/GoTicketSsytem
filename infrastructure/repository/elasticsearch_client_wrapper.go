// repository/elasticsearch_client_wrapper.go
package repository

import "github.com/olivere/elastic/v7"

// ElasticsearchClientWrapper wraps *elastic.Client to implement ElasticsearchClient.
type ElasticsearchClientWrapper struct {
	client *elastic.Client
}

// NewElasticsearchClientWrapper creates a new wrapper around *elastic.Client.
func NewElasticsearchClientWrapper(client *elastic.Client) *ElasticsearchClientWrapper {
	return &ElasticsearchClientWrapper{client: client}
}

// Index delegates to *elastic.Client's Index method.
func (w *ElasticsearchClientWrapper) Index() *elastic.IndexService {
	return w.client.Index()
}

// Search delegates to *elastic.Client's Search method.
func (w *ElasticsearchClientWrapper) Search() *elastic.SearchService {
	return w.client.Search()
}

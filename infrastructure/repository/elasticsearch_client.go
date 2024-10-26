package repository

import "github.com/olivere/elastic/v7"

// ElasticsearchClient defines an interface for the required Elasticsearch client methods
type ElasticsearchClient interface {
	Search() *elastic.SearchService
	Index() *elastic.IndexService
}

// repository/elasticsearch_ticket_repository.go
package repository

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"ticketing-system/application/dto"
	"ticketing-system/domain/entities"

	"github.com/olivere/elastic/v7"
)

// ElasticsearchTicketRepository uses an ElasticsearchClient for ticket data.
type ElasticsearchTicketRepository struct {
	esClient ElasticsearchClient
}

// NewElasticsearchTicketRepository creates a new ElasticsearchTicketRepository.
func NewElasticsearchTicketRepository(esClient ElasticsearchClient) *ElasticsearchTicketRepository {
	return &ElasticsearchTicketRepository{esClient: esClient}
}

// SaveToElasticsearch saves a ticket to the Elasticsearch index
func (r *ElasticsearchTicketRepository) SaveToElasticsearch(t *entities.Ticket) error {
	log.Printf("Indexing ticket with ID %d in Elasticsearch", t.ID)

	docID := strconv.Itoa(int(t.ID))

	_, err := r.esClient.Index().
		Index("tickets").
		Id(docID).
		BodyJson(t).
		Do(context.Background())
	if err != nil {
		log.Printf("Failed to index ticket in Elasticsearch: %v", err)
		return err
	}

	log.Printf("Successfully indexed ticket with ID %d in Elasticsearch", t.ID)
	return nil
}

// Search queries Elasticsearch for tickets based on the provided criteria
func (r *ElasticsearchTicketRepository) Search(query map[string]interface{}) ([]entities.Ticket, error) {
	// Build the query
	esQuery := elastic.NewBoolQuery()
	for key, value := range query {
		esQuery = esQuery.Must(elastic.NewMatchQuery(key, value))
	}

	// Perform the search
	searchResult, err := r.esClient.Search().
		Index("tickets").
		Query(esQuery).
		Do(context.Background())
	if err != nil {
		log.Printf("Failed to search tickets in Elasticsearch: %v", err)
		return nil, err
	}

	var tickets []entities.Ticket
	for _, hit := range searchResult.Hits.Hits {
		var t entities.Ticket
		err := json.Unmarshal(hit.Source, &t)
		if err != nil {
			log.Printf("Failed to unmarshal ticket: %v", err)
			continue
		}
		tickets = append(tickets, t)
	}

	return tickets, nil
}

// GetTickets retrieves tickets from Elasticsearch with filtering, sorting, and pagination
func (r *ElasticsearchTicketRepository) GetTickets(filters *dto.FilterOptions, sort dto.SortOptions, pageSize, page int) ([]entities.Ticket, int, error) {
	query := elastic.NewBoolQuery()

	// Apply filtering based on `created_at`
	if filters != nil {
		switch filters.FilterType {
		case "before":
			query = query.Filter(elastic.NewRangeQuery("CreatedAt").Lt(filters.FilterValue))
		case "after":
			query = query.Filter(elastic.NewRangeQuery("CreatedAt").Gt(filters.FilterValue))
		case "between":
			dateRange := strings.Split(filters.FilterValue, ",")
			if len(dateRange) == 2 {
				query = query.Filter(elastic.NewRangeQuery("CreatedAt").Gte(dateRange[0]).Lte(dateRange[1]))
			} else {
				log.Println("Invalid date range format for 'between' filter")
			}
		default:
			log.Println("Unknown filter type:", filters.FilterType)
		}
	}

	// Sorting logic
	sortField := "CreatedAt" // Default sort field
	if sort.SortName == "user_id" {
		sortField = "UserID"
	}
	ascending := sort.SortDir == "asc"
	searchService := r.esClient.Search().
		Index("tickets").
		Query(query).
		Sort(sortField, ascending)

	// Pagination logic
	searchService = searchService.From((page - 1) * pageSize).Size(pageSize)

	// Execute the search
	searchResult, err := searchService.Do(context.Background())
	if err != nil {
		log.Printf("Failed to search tickets in Elasticsearch: %v", err)
		return nil, 0, err
	}

	// Parse results
	var tickets []entities.Ticket
	for _, hit := range searchResult.Hits.Hits {
		var ticket entities.Ticket
		err := json.Unmarshal(hit.Source, &ticket)
		if err != nil {
			log.Printf("Failed to unmarshal ticket: %v", err)
			continue
		}
		tickets = append(tickets, ticket)
	}

	// Return the tickets, total count of hits, and no error
	return tickets, int(searchResult.TotalHits()), nil
}

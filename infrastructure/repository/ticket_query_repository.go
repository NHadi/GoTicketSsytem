package repository

import (
	"ticketing-system/application/dto"
	"ticketing-system/domain/entities"
)

// TicketQueryRepository defines the interface for querying the read model (e.g., Elasticsearch)
type TicketQueryRepository interface {
	// SaveToElasticsearch saves a ticket to the Elasticsearch index
	SaveToElasticsearch(t *entities.Ticket) error

	// Search is a generic search function that can be used to search tickets
	Search(query map[string]interface{}) ([]entities.Ticket, error)

	// GetTickets retrieves tickets with filters, sorting, and pagination
	GetTickets(filters *dto.FilterOptions, sort dto.SortOptions, pageSize, page int) ([]entities.Ticket, int, error)
}

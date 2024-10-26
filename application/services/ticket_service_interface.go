// application/services/ticket_service_interface.go

package services

import (
	"ticketing-system/application/dto"
	"ticketing-system/domain/entities"
)

// TicketServiceInterface defines the contract for TicketService
type TicketServiceInterface interface {
	CreateTicket(title, message string, userID uint) (*entities.Ticket, error)
	UpdateTicketStatus(id string, newStatus entities.TicketStatus) error
	GetTickets(filters *dto.FilterOptions, sort dto.SortOptions, pageSize, page int) ([]entities.Ticket, int, error)
}

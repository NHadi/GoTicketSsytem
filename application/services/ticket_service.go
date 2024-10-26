package services

import (
	"log"
	"strconv"
	"ticketing-system/application/dto"
	"ticketing-system/domain/entities"
	"ticketing-system/infrastructure/events"
	"ticketing-system/infrastructure/repository"
)

// TicketService handles ticket-related business logic
type TicketService struct {
	repo      repository.TicketRepository
	queryRepo repository.TicketQueryRepository
	publisher events.EventPublisher
}

// NewTicketService creates a new TicketService
func NewTicketService(repo repository.TicketRepository, publisher events.EventPublisher, queryRepo repository.TicketQueryRepository) *TicketService {
	return &TicketService{
		repo:      repo,
		queryRepo: queryRepo,
		publisher: publisher,
	}
}

// CreateTicket creates a new ticket and publishes the entities.created event
func (s *TicketService) CreateTicket(title, message string, userID uint) (*entities.Ticket, error) {
	log.Println("Creating new ticket...")

	// Create ticket using domain logic
	newTicket, err := entities.NewTicket(title, message, userID)
	if err != nil {
		log.Printf("Failed to create ticket entity: %v", err)
		return nil, err
	}
	log.Printf("Created ticket entity: %+v", newTicket)

	// Persist the new ticket
	err = s.repo.Save(newTicket)
	if err != nil {
		log.Printf("Failed to save ticket to the repository: %v", err)
		return nil, err
	}
	log.Printf("Ticket saved successfully with ID: %d", newTicket.ID)

	// Publish entities.created event
	err = s.publisher.PublishTicketCreated(newTicket)
	if err != nil {
		log.Printf("Failed to publish ticket.created event: %v", err)
		return nil, err
	}
	log.Println("ticket.created event published successfully")

	return newTicket, nil
}

// UpdateTicketStatus updates ticket status and publishes entities.status.updated event
func (s *TicketService) UpdateTicketStatus(id string, newStatus entities.TicketStatus) error {
	log.Printf("Updating ticket status for ID: %s", id)

	// Convert string id to uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Printf("Failed to parse ticket ID: %v", err)
		return err
	}

	// Find the ticket by its ID
	t, err := s.repo.FindByID(uint(idUint))
	if err != nil {
		log.Printf("Failed to find ticket with ID %d: %v", idUint, err)
		return err
	}
	log.Printf("Found ticket: %+v", t)

	// Update ticket status
	t.ChangeStatus(newStatus)
	log.Printf("Changed ticket status to: %s", newStatus)

	// Persist status change
	err = s.repo.UpdateStatus(t)
	if err != nil {
		log.Printf("Failed to update ticket status in the repository: %v", err)
		return err
	}
	log.Println("Ticket status updated successfully in the repository")

	// Publish entities.status.updated event
	err = s.publisher.PublishTicketStatusUpdated(t)
	if err != nil {
		log.Printf("Failed to publish ticket.status.updated event: %v", err)
		return err
	}
	log.Println("ticket.status.updated event published successfully")

	return nil
}

// GetTickets retrieves a list of tickets with filtering, sorting, and pagination
func (s *TicketService) GetTickets(filters *dto.FilterOptions, sort dto.SortOptions, pageSize, page int) ([]entities.Ticket, int, error) {
	// Validate and adjust pageSize to the nearest allowed value
	if pageSize != 10 && pageSize != 20 && pageSize != 30 && pageSize != 40 && pageSize != 50 {
		if pageSize < 10 {
			pageSize = 10
		} else if pageSize < 20 {
			pageSize = 10
		} else if pageSize < 30 {
			pageSize = 20
		} else if pageSize < 40 {
			pageSize = 30
		} else {
			pageSize = 40
		}
	}
	if page < 1 {
		page = 1
	}

	// Delegate to the query repository to get tickets
	tickets, totalCount, err := s.queryRepo.GetTickets(filters, sort, pageSize, page)
	if err != nil {
		return nil, 0, err
	}

	return tickets, totalCount, nil
}

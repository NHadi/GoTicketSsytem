package repository

import (
	"errors"
	"ticketing-system/domain/entities"

	"gorm.io/gorm"
)

// TicketRepository defines the interface for working with the Ticket persistence layer (SQL Server)
type TicketRepository interface {
	Save(t *entities.Ticket) error              // Save a new ticket to the database
	FindByID(id uint) (*entities.Ticket, error) // Find a ticket by its ID
	UpdateStatus(t *entities.Ticket) error      // Update the status of an existing ticket
}

// GormTicketRepository implements TicketRepository using GORM for SQL Server
type GormTicketRepository struct {
	db *gorm.DB
}

// NewGormTicketRepository creates a new instance of GormTicketRepository
func NewGormTicketRepository(db *gorm.DB) *GormTicketRepository {
	return &GormTicketRepository{db: db}
}

// Save persists a new ticket in the SQL Server database
func (r *GormTicketRepository) Save(t *entities.Ticket) error {
	// Perform the "Create" operation using GORM
	if err := r.db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

// FindByID retrieves a ticket by its ID from the SQL Server database
func (r *GormTicketRepository) FindByID(id uint) (*entities.Ticket, error) {
	var foundTicket entities.Ticket

	// Use GORM to find the ticket by its primary key (ID)
	if err := r.db.First(&foundTicket, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entities.ErrTicketNotFound
		}
		return nil, err
	}

	return &foundTicket, nil
}

// UpdateStatus updates the status of an existing ticket in the SQL Server database
func (r *GormTicketRepository) UpdateStatus(t *entities.Ticket) error {
	// Use GORM to update only the "status" column
	if err := r.db.Model(t).Update("status", t.Status).Error; err != nil {
		return err
	}
	return nil
}

package entities

import "time"

// TicketStatus defines valid ticket statuses
type TicketStatus string

const (
	Open     TicketStatus = "opn"
	Closed   TicketStatus = "cld"
	Assigned TicketStatus = "asn"
)

// Ticket represents the core entity for a support ticket
type Ticket struct {
	ID        uint
	Title     string
	Message   string
	UserID    uint
	Status    TicketStatus
	CreatedAt time.Time
}

// NewTicket is a factory method to create a new ticket
func NewTicket(title, message string, userID uint) (*Ticket, error) {
	if len(title) < 10 || len(title) > 100 {
		return nil, ErrInvalidTitle
	}
	if len(message) < 100 {
		return nil, ErrInvalidMessage
	}

	return &Ticket{
		Title:     title,
		Message:   message,
		UserID:    userID,
		Status:    Open,
		CreatedAt: time.Now(),
	}, nil
}

// ChangeStatus updates the ticket's status
func (t *Ticket) ChangeStatus(status TicketStatus) {
	t.Status = status
}

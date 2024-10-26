package responses

import "time"

// TicketResponse represents a single ticket in the response
type TicketResponse struct {
	TicketName   string    `json:"ticket_name"`
	TicketStatus string    `json:"ticket_status"`
	CreatedAt    time.Time `json:"created_at"`
	UserID       uint      `json:"user_id"`
}

// GetTicketListResponse represents the paginated response for the ticket list
type GetTicketListResponse struct {
	Tickets     []TicketResponse `json:"tickets"`
	TotalCount  int              `json:"total_count"`
	PageSize    int              `json:"page_size"`
	CurrentPage int              `json:"current_page"`
}

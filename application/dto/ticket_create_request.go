package dto

// CreateTicketRequest defines the request payload for creating a new ticket
type CreateTicketRequest struct {
	Title   string `json:"title" binding:"required,min=10,max=100"`
	Message string `json:"message" binding:"required,min=100"`
	UserID  uint   `json:"user_id" binding:"required"`
}

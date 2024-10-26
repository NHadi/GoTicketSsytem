package dto

// UpdateTicketStatusRequest defines the request payload for updating the status of an existing ticket
type UpdateTicketStatusRequest struct {
	Status string `json:"status" binding:"required,oneof='opn' 'cld' 'asn'"` // Limit to valid statuses
}

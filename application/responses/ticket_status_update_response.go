package responses

// TicketStatusUpdateResponse defines the response structure for a successful ticket status update
type TicketStatusUpdateResponse struct {
	Message   string `json:"message"`
	TicketID  string `json:"ticketID"`
	NewStatus string `json:"newStatus"`
}

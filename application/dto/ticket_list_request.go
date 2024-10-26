package dto

// GetTicketListRequest defines the request payload for getting the ticket list
type GetTicketListRequest struct {
	Filter   *FilterOptions `json:"filter,omitempty"`
	Sort     SortOptions    `json:"sort"`
	PageSize int            `json:"page_size"`
	Page     int            `json:"page"`
}

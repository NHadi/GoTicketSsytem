package api

import (
	"net/http"
	"ticketing-system/application/dto"
	"ticketing-system/application/responses"
	"ticketing-system/application/services"

	"github.com/gin-gonic/gin"
)

// QueryHandler handles HTTP requests related to querying tickets
type QueryHandler struct {
	service *services.TicketService
}

// NewQueryHandler creates a new QueryHandler with the given TicketService
func NewQueryHandler(service *services.TicketService) *QueryHandler {
	return &QueryHandler{service: service}
}

// GetTicketList handles retrieving a list of tickets with filtering, sorting, and pagination
// @Summary Get list of tickets
// @Description Retrieves a list of tickets with optional filtering, sorting, and pagination
// @Tags Tickets
// @Accept json
// @Produce json
// @Param request body dto.GetTicketListRequest true "Request payload for getting ticket list"
// @Success 200 {object} responses.GetTicketListResponse "Successful response with ticket list"
// @Failure 400 {object} responses.ErrorResponse "Invalid request payload"
// @Failure 500 {object} responses.ErrorResponse "Failed to retrieve tickets"
// @Router /tickets/list [post]
func (h *QueryHandler) GetTicketList(c *gin.Context) {
	var req dto.GetTicketListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	// Fetch tickets using the service layer
	tickets, totalCount, err := h.service.GetTickets(req.Filter, req.Sort, req.PageSize, req.Page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Failed to retrieve tickets", Details: err.Error()})
		return
	}

	// Map to response model
	var ticketResponses []responses.TicketResponse
	for _, ticket := range tickets {
		ticketResponses = append(ticketResponses, responses.TicketResponse{
			TicketName:   ticket.Title,
			TicketStatus: string(ticket.Status),
			CreatedAt:    ticket.CreatedAt,
			UserID:       ticket.UserID,
		})
	}

	// Return paginated response
	c.JSON(http.StatusOK, responses.GetTicketListResponse{
		Tickets:     ticketResponses,
		TotalCount:  totalCount,
		PageSize:    req.PageSize,
		CurrentPage: req.Page,
	})
}

package api

import (
	"log"
	"net/http"
	"ticketing-system/application/dto"
	"ticketing-system/application/responses"
	"ticketing-system/application/services"
	"ticketing-system/domain/entities"

	"github.com/gin-gonic/gin"
)

// CommandHandler defines the structure for handling ticket-related commands
type CommandHandler struct {
	service services.TicketServiceInterface
}

// NewCommandHandler creates a new instance of CommandHandler
func NewCommandHandler(service services.TicketServiceInterface) *CommandHandler {
	return &CommandHandler{service: service}
}

// @Summary Create a new ticket
// @Description Create a new ticket in the system
// @Tags Tickets
// @Accept json
// @Produce json
// @Param ticket body CreateTicketRequest true "Ticket Data"
// @Success 201 {object} entities.Ticket
// @Failure 400 {object} responses.ErrorResponse "Invalid input data"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /tickets [post]
func (h *CommandHandler) CreateTicket(c *gin.Context) {
	var req dto.CreateTicketRequest

	// Bind and validate the JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	// Call the service layer to create the ticket
	newTicket, err := h.service.CreateTicket(req.Title, req.Message, req.UserID)
	if err != nil {
		log.Printf("Failed to create ticket: %v", err)
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error:   "Failed to create ticket",
			Details: err.Error(),
		})
		return
	}

	// Return the newly created ticket in the response with 201 Created status
	c.JSON(http.StatusCreated, newTicket)
}

// @Summary Update ticket status
// @Description Update the status of an existing ticket
// @Tags Tickets
// @Accept json
// @Produce json
// @Param id path string true "Ticket ID"
// @Param status body UpdateTicketStatusRequest true "New Status"
// @Success 200 {object} TicketStatusUpdateResponse
// @Failure 400 {object} responses.ErrorResponse "Invalid status"
// @Failure 500 {object} responses.ErrorResponse "Failed to update ticket status"
// @Router /tickets/{id}/status [patch]
func (h *CommandHandler) UpdateTicketStatus(c *gin.Context) {
	var req dto.UpdateTicketStatusRequest

	// Bind and validate the JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	// Parse the ticket ID from the URL path parameter
	ticketID := c.Param("id")

	// Map the string status from the request to the domain's TicketStatus type
	var newStatus entities.TicketStatus
	switch req.Status {
	case "opn":
		newStatus = entities.Open
	case "cld":
		newStatus = entities.Closed
	case "asn":
		newStatus = entities.Assigned
	default:
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid status"})
		return
	}

	// Call the service layer to update the ticket status
	err := h.service.UpdateTicketStatus(ticketID, newStatus)
	if err != nil {
		log.Printf("Failed to update ticket status: %v", err)
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error:   "Failed to update ticket status",
			Details: err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, responses.TicketStatusUpdateResponse{
		Message:   "Ticket status updated successfully",
		TicketID:  ticketID,
		NewStatus: string(newStatus),
	})
}

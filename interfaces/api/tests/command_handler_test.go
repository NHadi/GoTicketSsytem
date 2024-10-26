package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"ticketing-system/application/dto"
	"ticketing-system/application/responses"
	"ticketing-system/application/services/mocks"
	"ticketing-system/domain/entities"
	"ticketing-system/interfaces/api"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTicket(t *testing.T) {
	// Initialize the mock service
	mockService := new(mocks.TicketServiceMock)
	handler := api.NewCommandHandler(mockService)

	// Define request payload and expected response
	requestPayload := dto.CreateTicketRequest{
		Title:   "Sample Ticket",
		Message: "This is a test message",
		UserID:  1,
	}
	expectedTicket := &entities.Ticket{
		ID:      1,
		Title:   "Sample Ticket",
		Message: "This is a test message",
		UserID:  1,
		Status:  entities.Open,
	}

	// Set up mock behavior
	mockService.On("CreateTicket", requestPayload.Title, requestPayload.Message, requestPayload.UserID).
		Return(expectedTicket, nil)

	// Create a new Gin context and recorder
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Prepare the JSON request
	requestBody, _ := json.Marshal(requestPayload)
	c.Request, _ = http.NewRequest("POST", "/tickets", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the CreateTicket handler
	handler.CreateTicket(c)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	var response entities.Ticket
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, expectedTicket.ID, response.ID)
	assert.Equal(t, expectedTicket.Title, response.Title)
	mockService.AssertExpectations(t)
}

func TestUpdateTicketStatus(t *testing.T) {
	// Initialize the mock service
	mockService := new(mocks.TicketServiceMock)
	handler := api.NewCommandHandler(mockService)

	// Define request payload and path parameter
	ticketID := "1"
	requestPayload := dto.UpdateTicketStatusRequest{
		Status: "cld",
	}
	expectedStatus := entities.Closed

	// Set up mock behavior
	mockService.On("UpdateTicketStatus", ticketID, expectedStatus).
		Return(nil)

	// Create a new Gin context and recorder
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Prepare the JSON request
	requestBody, _ := json.Marshal(requestPayload)
	c.Request, _ = http.NewRequest("PATCH", "/tickets/"+ticketID+"/status", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Set the ticket ID parameter
	c.Params = gin.Params{{Key: "id", Value: ticketID}}

	// Call the UpdateTicketStatus handler
	handler.UpdateTicketStatus(c)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response responses.TicketStatusUpdateResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Ticket status updated successfully", response.Message)
	assert.Equal(t, ticketID, response.TicketID)
	assert.Equal(t, string(expectedStatus), response.NewStatus)

	mockService.AssertExpectations(t)
}

package entities

import "errors"

// Errors related to the Ticket domain
var (
	// ErrInvalidTitle is returned when the ticket title does not meet the business rules
	ErrInvalidTitle = errors.New("ticket title must be between 10 and 100 characters")

	// ErrInvalidMessage is returned when the ticket message does not meet the business rules
	ErrInvalidMessage = errors.New("ticket message must be at least 100 characters long")

	// ErrInvalidStatus is returned when an invalid status is provided
	ErrInvalidStatus = errors.New("invalid ticket status")

	// ErrTicketNotFound is returned when a ticket is not found in the repository
	ErrTicketNotFound = errors.New("ticket not found")
)

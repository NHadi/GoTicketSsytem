
# Ticketing System API

This repository contains a microservices-based ticketing system with support for CRUD operations, filtering, sorting, and event-driven communication using RabbitMQ and Elasticsearch.

## Architecture Overview

The system follows a Domain-Driven Design (DDD) approach, employing:
- **Golang** for backend services
- **Elasticsearch** for search and storage of ticket queries
- **RabbitMQ** for handling events
- **SQL Server** for primary storage
- **Swagger** for API documentation

## Project Structure

- **application**: Contains service logic and DTOs for managing ticket operations.
- **domain**: Defines entities for the system (e.g., Ticket).
- **infrastructure**: Manages external connections, including repositories and event handling.
- **interfaces**: Contains API handlers and request validation.

## Key Features

- **Create Ticket API**: Allows support agents to create new tickets.
- **Get Ticket List API**: Supports filtering, sorting, and pagination.
- **Event-Driven Updates**: Ticket updates are published and consumed via RabbitMQ for Elasticsearch indexing.

## Setup Instructions

### Prerequisites
- Go 1.23+
- Docker and Docker Compose
- SQL Server, Elasticsearch, RabbitMQ

### Steps
1. **Build and Run**: Use Docker Compose to start all services.
   ```sh
   docker-compose up --build
   ```

2. **Swagger Documentation**: Available at `http://localhost:8080/swagger/index.html`.

## Testing

### Unit Tests
Run unit tests using:
```sh
go test ./...
```

### Postman Collection
A Postman collection `postman_collection.json` is included in the repository for testing.

## Endpoints

### Create Ticket
- **POST** `/tickets`
- **Request Body**:
  ```json
  {
    "title": "Sample Title",
    "message": "Sample ticket message exceeding 100 characters...",
    "user_id": 1
  }
  ```

### Get Ticket List
- **POST** `/tickets/list`
- **Request Body**:
  ```json
  {
    "filter": {
      "filter_name": "created_at",
      "filter_type": "before",
      "filter_value": "2022-11-14"
    },
    "sort": {
      "sort_name": "created_at",
      "sort_dir": "asc"
    },
    "page_size": 10
  }
  ```

## License

This project is licensed under the MIT License.

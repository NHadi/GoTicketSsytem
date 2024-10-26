package main

import (
	"log"
	"os"
	"ticketing-system/application/services"
	"ticketing-system/domain/entities"
	"ticketing-system/infrastructure/events"
	"ticketing-system/infrastructure/repository"
	"ticketing-system/interfaces/api"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/streadway/amqp"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	_ "ticketing-system/docs" // Import your generated docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Ticketing System API
// @version 1.0
// @description This is a sample server for a ticketing system.
// @host localhost:8080
// @BasePath /

// connectToSQLServer connects to SQL Server with retry logic.
func connectToSQLServer(dbURL string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ { // Retry up to 10 times
		db, err = gorm.Open(sqlserver.Open(dbURL), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		log.Printf("Failed to connect to SQL Server. Retrying... (%d/10)\n", i+1)
		time.Sleep(5 * time.Second)
	}
	return nil, err
}

func main() {
	// Load environment variables
	dbURL := os.Getenv("SQLSERVER_URL")
	elasticURL := os.Getenv("ELASTICSEARCH_URL")
	rabbitMQURL := os.Getenv("RABBITMQ_URL")

	// Connect to SQL Server
	db, err := connectToSQLServer(dbURL)
	if err != nil {
		log.Fatal("Failed to connect to SQL Server:", err)
	}

	// Migrate the database schema for command operations
	db.AutoMigrate(&entities.Ticket{})

	// Connect to Elasticsearch for query operations
	esClient, err := elastic.NewClient(elastic.SetURL(elasticURL))
	if err != nil {
		log.Fatal("Failed to connect to Elasticsearch:", err)
	}

	// Connect to RabbitMQ for event handling
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a RabbitMQ channel:", err)
	}
	defer channel.Close()

	// Initialize repositories
	ticketRepo := repository.NewGormTicketRepository(db) // For command operations

	esClientWrapper := repository.NewElasticsearchClientWrapper(esClient)
	ticketQueryRepo := repository.NewElasticsearchTicketRepository(esClientWrapper) // For query operations

	// Initialize services
	eventPublisher := events.NewEventPublisher(channel)
	ticketService := services.NewTicketService(ticketRepo, eventPublisher, ticketQueryRepo)

	// Initialize the consumer for syncing events to Elasticsearch
	eventConsumer := events.NewEventConsumer(ticketQueryRepo, channel)
	eventConsumer.StartConsumer()

	// Ensure graceful shutdown of the consumer
	defer func() {
		if err := eventConsumer.StopConsumer(); err != nil {
			log.Printf("Failed to stop consumer: %v", err)
		}
	}()

	// Initialize API handlers
	r := gin.Default()
	commandHandler := api.NewCommandHandler(ticketService)
	queryHandler := api.NewQueryHandler(ticketService)

	// Add Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/tickets", commandHandler.CreateTicket)
	r.PATCH("/tickets/:id/status", commandHandler.UpdateTicketStatus)
	r.POST("/tickets/list", queryHandler.GetTicketList)

	// Start the HTTP server
	log.Println("Starting combined command and query service on port 8080")
	r.Run(":8080")
}

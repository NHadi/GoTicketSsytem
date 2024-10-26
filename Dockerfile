# Build stage
FROM golang:1.23 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application as a statically linked binary for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ticketing-service main.go

# Final stage: use a minimal image to run the binary
FROM alpine:latest
WORKDIR /root/

# Copy the statically built binary from the builder stage
COPY --from=builder /app/ticketing-service .

# Ensure the binary is executable
RUN chmod +x /root/ticketing-service

# Expose the application's port
EXPOSE 8080

# Run the application
CMD ["/root/ticketing-service"]

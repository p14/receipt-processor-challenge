# Use the official Golang image as the build environment
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o receipt-processor ./cmd/receipt-processor

# Use a minimal image for the final build
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/receipt-processor .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./receipt-processor"]

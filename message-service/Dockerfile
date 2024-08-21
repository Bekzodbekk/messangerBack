# Base image
FROM golang:1.22.5-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o message-service .

# Expose the port
EXPOSE 8002

# Command to run the application
CMD ["./message-service"]

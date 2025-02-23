# Use an official Golang image for building the application
FROM golang:1.23.1 AS builder

# Set the working directory in the builder container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the application from cmd/main.go, outputting a binary named 'app'
RUN go build -o app ./cmd

# Use a minimal base image for the final container
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/app /app

# Expose the application's port (e.g., 8080)
EXPOSE 8080

# Set the command to run the application
CMD ["/app"]

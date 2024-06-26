# Start from the official Golang image
FROM golang:1.16-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main cmd/main.go

# Start a new stage from scratch
FROM alpine:latest

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Copy the config file
COPY config.json /app/config.json

# Expose port 8080 (if your application serves HTTP)
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]

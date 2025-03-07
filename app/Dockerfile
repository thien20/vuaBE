# Stage 1: Build
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# COPY config/config.json /app/config/config.json


# Build the Go app
RUN go build -o main .

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /app/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./main"]
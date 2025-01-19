# Use an official Golang runtime as a parent image
FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any needed Go modules
RUN go mod tidy

# Build the Go app
# RUN go build -o main .

# Run the Go app
CMD ["go", "run", "."]

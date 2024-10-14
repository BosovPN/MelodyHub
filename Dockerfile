# Use the official Golang image for building the application
FROM golang:1.23.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o melodyhub ./cmd/main.go

# Use a lightweight base image for the final image
FROM alpine:latest

# Set the working directory for the final image
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/melodyhub .

# Copy the .env file to the container (optional)
COPY .env ./

# Expose a port if necessary (e.g., port 8080)
EXPOSE 8080

# Command to run the application
CMD ["./melodyhub"]
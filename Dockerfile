# Stage 1: Build the Go binary
FROM golang:1.22.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire Go project into the container
COPY . .

# Build the Go binary
RUN go build -o main .

# Stage 2: Create a minimal final image with just the Go binary
FROM debian:bullseye-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 (or whatever port your app uses)
EXPOSE 8081

# Run the Go binary
CMD ["./main"]

# Stage 1: Build the Go binary
FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go binary as a static executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Use a minimal image like Alpine since the binary is statically linked
FROM alpine:latest

WORKDIR /app

# Copy the config.json file
COPY config.json .

COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x ./main

EXPOSE 8080

CMD ["./main"]
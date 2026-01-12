# --- Stage 1: Build ---
FROM golang:1.25.3-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Install git (needed for Go modules)
RUN apk add --no-cache git

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o muchtodo ./cmd/api

# --- Stage 2: Run ---
FROM alpine:latest

# Set working directory inside container
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/muchtodo .

# Copy .env if you want it inside container (optional)
COPY .env .env

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./muchtodo"]

# Stage 1: Build the Go binary
FROM golang:1.25.3-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Install git and other necessary tools
RUN apk add --no-cache git bash

# Copy go.mod and go.sum first (for caching dependencies)
COPY app/Server/MuchToDo/go.mod app/Server/MuchToDo/go.sum ./

# Download dependencies
RUN go mod download

# Copy all source code
COPY app/Server/MuchToDo/. .

# Ensure external modules needed for your code are fetched
RUN go get github.com/golang-jwt/jwt/v5 \
    github.com/spf13/viper \
    github.com/redis/go-redis/v9 \
    go.mongodb.org/mongo-driver/mongo \
    go.mongodb.org/mongo-driver/mongo/options \
    go.mongodb.org/mongo-driver/mongo/readpref \
    go.mongodb.org/mongo-driver/bson \
    github.com/swaggo/files \
    github.com/swaggo/gin-swagger \
    golang.org/x/crypto/pbkdf2

# Tidy up modules
RUN go mod tidy

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o muchtodo ./cmd/api

# Stage 2: Create a small runtime image
FROM alpine:3.22

# Set working directory inside container
WORKDIR /app

# Copy the compiled binary from the builder
COPY --from=builder /app/muchtodo .

# Expose the port your app listens on
EXPOSE 8080

# Run the application
CMD ["./muchtodo"]

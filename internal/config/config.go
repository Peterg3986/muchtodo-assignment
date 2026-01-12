package config

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectMongo reads the MongoDB URI from environment variable and connects
func ConnectMongo() *mongo.Client {
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI environment variable not set")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Ping MongoDB to ensure connection is alive
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    log.Println("✅ Connected to MongoDB successfully")
    return client
}


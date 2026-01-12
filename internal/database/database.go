package database

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectMongo connects to MongoDB running in Docker.
func ConnectMongo() (*mongo.Client, error) {
    // Hardcoded MongoDB URI for Docker container
    uri := "mongodb://admin:admin123@muchtodo-mongo:27017/muchtodo?authSource=admin"

    // Set a timeout for the connection attempt.
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }

    // Ping the primary node to verify that the connection is alive.
    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        return nil, err
    }

    return client, nil
}


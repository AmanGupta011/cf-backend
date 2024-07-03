package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func DbConnection(client *mongo.client, err error){
// 	err = godotenv.Load()
// 	if err != nil{
// 		log.Fatal("Error loading .env file")
// 	}
// 	uri := os.Getenv("MONGODB_URI")
// 	clientOptions := options.Client().ApplyURI(uri)
// 	ctx := context.TODO()
// 	client, err = mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return
// }

func DbConnection() (*mongo.Client, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the MongoDB URI from the environment variables
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Create a context
	ctx := context.TODO()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Return the MongoDB client and nil error
	return client, nil
}

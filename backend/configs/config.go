package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(EnvMongoURI())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB")

	DB = client
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		err := GetMongoClient()
		if err != nil {
			log.Fatalf("Failed to get MongoDB client: %v", err)
		}
	}
	collection := DB.Database(EnvMongoDatabaseName()).Collection(collectionName)
	return collection
}

// DB Client instance
var DB *mongo.Client

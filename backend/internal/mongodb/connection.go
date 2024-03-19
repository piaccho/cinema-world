package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// // Collection names (corresponding to JSON files)
// var colNames = map[string]string{
// 	"categories.json": "categories",
// 	"genres.json":     "genres",
// 	"halls.json":      "halls",
// 	"movies.json":     "movies",
// 	"showing.json":    "showings",
// 	"users.json":      "users",
// }

var client *mongo.Client

func connectToMongo(ctx context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}
	log.Println("Connected to MongoDB")
	return client, nil
}

// GetMongoClient returns the established MongoDB client
func GetMongoClient() (*mongo.Client, error) {
	if client == nil {
		uri := os.Getenv("MONGODB_URI")
		var err error
		client, err = connectToMongo(context.Background(), uri)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

// func createDatabase(client *mongo.Client, dbName string) error {
// 	command := bson.D{{Key: "use", Value: dbName}}
// 	err := client.Database(dbName).RunCommand(context.Background(), command).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func DropDatabase(client *mongo.Client, dbName string) {
// 	err := client.Database(dbName).Drop(context.Background())
// 	if err != nil {
// 		log.Printf("Error dropping database: %v\n", err)
// 	}
// 	log.Printf("Dropped database %s\n", dbName)
// }

// func InitializeDatabase(client *mongo.Client, dbName string) {
// 	// Create a new database
// 	err := createDatabase(client, dbName)
// 	if err != nil {
// 		log.Printf("Error creating database: %v\n", err)
// 	}

// 	// Create collections and import data on startup
// 	for filePath, collectionName := range colNames {
// 		// Construct the mongoimport command
// 		cmd := exec.Command("mongoimport", "--uri", os.Getenv("MONGODB_URI"), "--db", dbName, "--collection", collectionName, "--file", filePath, "--jsonArray")

// 		// Run the command
// 		err := cmd.Run()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// err := createCollectionAndImportData(client, fileName, colName)
// 		// if err != nil {
// 		// 	log.Printf("Error creating collection %s or importing data: %v\n", colName, err)
// 		// }
// 	}

// }

// func createCollectionAndImportData(client *mongo.Client, fileName string, colName string) error {
// 	ctx := context.Background()
// 	db := client.Database(os.Getenv("MONGODB_NAME"))

// 	// Get list of all collections
// 	collections, err := db.ListCollectionNames(ctx, bson.M{})
// 	if err != nil {
// 		return fmt.Errorf("error getting collection names: %v", err)
// 	}

// 	// Check if collection already exists
// 	for _, collection := range collections {
// 		if collection == colName {
// 			fmt.Printf("Collection %s already exists.\n", colName)
// 			return nil
// 		}
// 	}

// 	// Create collection
// 	err = db.CreateCollection(ctx, colName)
// 	if err != nil {
// 		return fmt.Errorf("error creating collection %s: %v", colName, err)
// 	}

// 	// Read data from JSON file
// 	filePath := fmt.Sprintf("init_data/%s", fileName)
// 	data, err := os.ReadFile(filePath)
// 	if err != nil {
// 		return fmt.Errorf("error reading file %s: %v", fileName, err)
// 	}

// 	// Unmarshal data into slice of appropriate type (adjust based on your data structure)
// 	var documents []interface{}
// 	err = json.Unmarshal(data, &documents)
// 	if err != nil {
// 		return fmt.Errorf("error unmarshalling JSON data from %s: %v", fileName, err)
// 	}

// 	// Insert documents into collection
// 	result, err := db.Collection(colName).InsertMany(ctx, documents)
// 	if err != nil {
// 		return fmt.Errorf("error inserting data into collection %s: %v", colName, err)
// 	}

// 	fmt.Printf("Successfully inserted %d documents into collection %s\n", len(result.InsertedIDs), colName)
// 	return nil
// }

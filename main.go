package KeDuBack

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func herror_handling(MongoURL *string, SECRET *string) int {
	error := godotenv.Load()
	if error != nil {
		fmt.Println("Error loading .env file")
		return 1
	}
	*MongoURL = os.Getenv("MONGO_URL")
	*SECRET = os.Getenv("SECRET")
	if *MongoURL == "" || *SECRET == "" {
		fmt.Println("Error loading .env file")
		return 1
	}
	return 0
}

func connectDB(MongoURL string) *mongo.Database {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(MongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB")
		return nil
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Error connecting to MongoDB")
		return nil
	}
	fmt.Println("Connected to MongoDB!")
	return client.Database("KeDuFront")
}

func startServer() {
	// Start the server
	return

}

func main() {
	var SECRET string
	var MongoURL string

	if herror_handling(&MongoURL, &SECRET) == 1 {
		return
	}
	fmt.Println("MongoURL: ", MongoURL)
	fmt.Println("SECRET: ", SECRET)
	my_db := connectDB(MongoURL)
	if my_db == nil {
		return
	}
	defer my_db.Client().Disconnect(context.Background())
	startServer()
	return
}

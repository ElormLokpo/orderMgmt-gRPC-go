package database

import (
	"context"
	"fmt"
	"notification-service/errHandler"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {
	err := godotenv.Load()
	errHandler.ErrHandler(err, "Error loading env files")

	mongo_url := os.Getenv("MONGO_URL")

	dbOptions := options.Client().ApplyURI(mongo_url)
	client, err = mongo.Connect(context.Background(), dbOptions)
	errHandler.ErrHandler(err, "Error connecting to mongo Database")

	fmt.Println("Database connected successfully")
}

func GetClient() *mongo.Client {
	return client
}

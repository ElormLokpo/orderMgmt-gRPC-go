package database

import (
	"context"
	"fmt"
	"os"

	"product-service/errhandler"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {
	err := godotenv.Load()
	errhandler.ErrHandler(err)

	mongo_url := os.Getenv("MONGO_URL")
	dbOptions := options.Client().ApplyURI(mongo_url)
	client, err = mongo.Connect(context.Background(), dbOptions)
	errhandler.ErrHandler(err)

	err = client.Ping(context.Background(), nil)
	errhandler.ErrHandler(err)

	fmt.Println("Connected to database successfully")

}

func GetClient() *mongo.Client {
	return client
}

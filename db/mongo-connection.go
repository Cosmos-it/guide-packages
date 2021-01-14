package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Conn ...
var Conn *mongo.Database

// connection ...
var connection string

var (
	dbName         string = os.Getenv("DATABASE_NAME")
	dbUsername     string = os.Getenv("DATABASE_USERNAME")
	dbPort         string = os.Getenv("DATABASE_HOST_PORT")
	dbPassword     string = os.Getenv("DATABASE_PASSWORD")
	devEnvironment string = os.Getenv("DEV_ENVIRONMENT") // Development environment is either dev/prod
)

// Connect ...
func Connect() {

	fmt.Println("Test", os.Getenv("DEV_ENVIRONMENT"))

	if devEnvironment == "dev" {
		connection = fmt.Sprintf("mongodb://localhost:27017")
	} else {
		connection = fmt.Sprintf("mongodb://%s:%s@%s/%s", dbUsername, dbPassword, dbPort, dbName)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(connection))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	print("\nConnected to MongoDB\n")
	Conn = client.Database(dbName)
}

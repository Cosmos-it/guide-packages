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

// Conn Global access to the database instance ...
var Conn *mongo.Database // Global variable

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

	if devEnvironment == "dev" {
		connection = fmt.Sprintf("mongodb://localhost:27017")
		clientConnect()

	} else if devEnvironment == "prod" {
		connection = fmt.Sprintf("mongodb://%s:%s@%s/%s", dbUsername, dbPassword, dbPort, dbName)
		clientConnect()

	} else {
		fmt.Println("No environment have been initialized")
	}

}

// [clientConnect] Connect to the database when environment variable is providedclienConnect ....
func clientConnect() {
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

	environment := fmt.Sprintf("\nConnected to Database in %s environment", devEnvironment)
	fmt.Println(environment)

	Conn = client.Database(dbName)
}

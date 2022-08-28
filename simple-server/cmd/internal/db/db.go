package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

const (

	DataBase = "product-api"
)

var clientInstance *mongo.Client 

var mongoOnce sync.Once 

var clientInstanceError error

func GetMongoClient()(*mongo.Client,error){
	mongoOnce.Do(func() {
		clientOption := options.Client().ApplyURI(goDotEnvVariable("DATABASE_URI"))

		client, err := mongo.Connect(context.TODO(),clientOption)

		clientInstance = client
		
		clientInstanceError = err 
	})

	return clientInstance ,clientInstanceError
}
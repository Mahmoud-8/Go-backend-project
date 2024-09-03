package main


import (
"fmt"
"log"
"os"
"github.com/gofiber/fiber/v2"
"context"
go get go.mongodb.org/mongo-driver/mongo
import "go.mongodb.org/mongo-driver/mongo/options"go get github.com/joho/godotenv
)
type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json: "completed"`
	Body      string `"json: "body"`
}
var collection *mongo.Collection

func main() {
fmt.Println("hello world") 
err := godotenv.Load(".env")
if err != nil {
	log.Fatal("Error loading .env file")
}

MONGODB_URI := os.Getenv("MONGODB_URI")
client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))

if err != nil {
	log.Fatal(err)
}

err := client.Ping(context.TODO(), nil)

}
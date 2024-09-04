package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

err = client.Ping(context.Background(), nil)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")


collection = client.Database("fiber-hrms").Collection("todos")
app := fiber.New()
// Get todos
app.Get("/api/todos", getTodos)
app.Post("/api/todos", createTodo)
app.Patch("/api/todos/:id", updateTodo)
app.Delete("/api/todos/:id", DeletTodo)

port := os.Getenv("PORT")
if port == "" {
	port = "5000"
}

log.Fatal(app.Listen("0.0.0.0:", port))
}

func getTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}
func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
}

if todo.Body == "" {
	return c.Status(400).JSON(fiber.Map{
		"error": "Body is required",
	})
}


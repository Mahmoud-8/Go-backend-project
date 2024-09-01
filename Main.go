package main


import (
"fmt"

"go.mongodb.org/mongo-driver/mongo"
)
type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json: "completed"`
	Body      string `"json: "body"`
}
var collection *mongo.Collection

func main() {
fmt.Println("hello world") I
}
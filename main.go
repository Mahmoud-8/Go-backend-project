package main

import (
        "fmt"
        "log"
)

func main() {
        app := fiber.New() 

        fmt.Println("Hello worlgfghd")
        log.Fatal(app.Listen(":4000"))
}
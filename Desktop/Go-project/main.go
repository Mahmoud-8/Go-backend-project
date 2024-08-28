package main

import (
        "fmt"
        "log"
        "github.com/gofiber/fiber/v2"
)

func main() {
        app := fiber.New() // Declare app here

        fmt.Println("Hello world")
        log.Fatal(app.Listen(":4000"))
}
package main

import (
        "fmt"
        "log"
        "github.com/gofiber/fiber/v2"
)

func main() {
        app := fiber.New() 

        app.Get("/", func(c *fiber.Ctx) error {
                return c. Status (200). JSON (fiber.Map{"msg": "hello world"})
                })

        fmt.Println("Hello worlgfghd")
        log.Fatal(app.Listen(":4000"))
}
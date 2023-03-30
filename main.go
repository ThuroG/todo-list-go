package main

import (
    "log"

    "main/config"
    "main/handlers"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Instantiate Fiber App
    app := fiber.New()

    // Connect to Database
    config.Connect()

    // Create Handlers
    app.Get("/todos", handlers.GetTodoItems)
    app.Get("/todos/:id", handlers.GetTodoItem)
    app.Post("/todos", handlers.AddTodoItem)
    app.Put("/todos/:id", handlers.UpdateTodoItem)
    app.Delete("/todos/:id", handlers.RemoveTodoItem)

    log.Fatal(app.Listen(":3000"))
}

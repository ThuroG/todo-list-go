package handlers

import (
    "main/config"
    "main/entities"
    "github.com/gofiber/fiber/v2"
)

func GetTodoItems(c *fiber.Ctx) error {
    var todos []entities.Todo

    config.Database.Find(&todos)
    return c.Status(200).JSON(todos)
}

func GetTodoItem(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo entities.Todo

    result := config.Database.Find(&todo, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.Status(200).JSON(&todo)
}

func AddTodoItem(c *fiber.Ctx) error {
    todo := new(entities.Todo)

    if err := c.BodyParser(todo); err != nil {
        return c.Status(503).SendString(err.Error())
    }

    config.Database.Create(&todo)
    return c.Status(201).JSON(todo)
}

func UpdateTodoItem(c *fiber.Ctx) error {
    todo := new(entities.Todo)
    id := c.Params("id")

    if err := c.BodyParser(todo); err != nil {
        return c.Status(503).SendString(err.Error())
    }

    config.Database.Where("id = ?", id).Updates(&todo)
    return c.Status(200).JSON(todo)
}

func RemoveTodoItem(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo entities.Todo

    result := config.Database.Delete(&todo, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.SendStatus(200)
}
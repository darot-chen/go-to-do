package handler

import (
	"github.com/darot-chen/go-to-do/databases"
	"github.com/darot-chen/go-to-do/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTodo(c *fiber.Ctx) error {
	db := databases.DB

	var todos []models.Todo

	db.Find(&todos)

	if len(todos) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "success", "data": todos})
}

func CreateTodo(c *fiber.Ctx) error {
	db := databases.DB

	todo := new(models.Todo)

	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Data invalid", "data": err})
	}

	todo.ID = uuid.New()

	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Cannot create todo", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success", "data": todo})
}

func UpdateTodo(c *fiber.Ctx) error {
	type UpdateTodo struct {
		Item      string `json:"item"`
		Completed int    `json:"completed"`
	}

	db := databases.DB

	todo := new(models.Todo)

	id := c.Params("id")

	db.Find(&todo, "id = ?", id)
	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Todo not found", "data": nil})
	}

	var updateTodo UpdateTodo
	err := c.BodyParser(&updateTodo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Data input invalid", "data": err})
	}

	todo.Item = updateTodo.Item
	todo.Completed = updateTodo.Completed

	db.Save(&todo)
	return c.JSON(fiber.Map{"status": "success", "message": "Success", "data": todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	db := databases.DB

	todo := new(models.Todo)
	id := c.Params("id")

	db.Find(&todo, "id = ?", id)

	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Todo not found", "data": nil})
	}

	err := db.Delete(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Cannot delete cannot", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success", "data": nil})
}

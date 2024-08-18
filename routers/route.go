package routers

import (
	"github.com/darot-chen/go-to-do/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app fiber.Router) {
	api := app.Group("/api")

	api.Get("/", handler.GetTodo)
	api.Post("/", handler.CreateTodo)
	api.Put("/:id", handler.UpdateTodo)
	api.Delete("/:id", handler.DeleteTodo)
}

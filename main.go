package main

import (
	"log"

	"github.com/darot-chen/go-to-do/databases"
	"github.com/darot-chen/go-to-do/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	databases.ConnectDB()

	routers.SetupRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("App is running")
	})

	log.Fatal(app.Listen(":3000"))
}

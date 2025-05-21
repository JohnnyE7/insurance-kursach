package main

import (
	"insurance-backend/config"
	"insurance-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.InitDB()

	app := fiber.New()

	api := app.Group("/api")

	api.Get("/clients", handlers.GetClients(db))
	api.Get("/clients/:id", handlers.GetClient(db))
	api.Post("/clients", handlers.CreateClient(db))
	api.Put("/clients/:id", handlers.UpdateClient(db))
	api.Delete("/clients/:id", handlers.DeleteClient(db))

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"insurance-backend/config"
	"insurance-backend/routes"
)

func main() {
	app := fiber.New()

	db := config.InitDB()
	routes.SetupClientRoutes(app, db)

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}

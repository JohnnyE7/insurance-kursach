package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"insurance-backend/handlers"
)

func SetupClientRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.NewClientHandler(db)

	app.Get("/clients", h.GetClients)
	app.Get("/clients/:id", h.GetClient)
	app.Post("/clients", h.CreateClient)
	app.Put("/clients/:id", h.UpdateClient)
	app.Delete("/clients/:id", h.DeleteClient)
}

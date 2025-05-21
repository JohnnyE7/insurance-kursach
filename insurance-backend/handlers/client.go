package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"insurance-backend/models"
)

func GetClients(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var clients []models.Client

		db.Find(&clients)
		return c.JSON(clients)
	}
}

func GetClient(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var client models.Client
		if err := db.First(&client, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Client not found"})
		}
		return c.JSON(client)
	}
}

func CreateClient(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		client := new(models.Client)
		if err := c.BodyParser(client); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}
		db.Create(&client)
		return c.Status(201).JSON(client)
	}
}

func UpdateClient(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var client models.Client
		if err := db.First(&client, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Client not found"})
		}
		if err := c.BodyParser(&client); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}
		db.Save(&client)
		return c.JSON(client)
	}
}

func DeleteClient(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		db.Delete(&models.Client{}, id)
		return c.SendStatus(204)
	}
}

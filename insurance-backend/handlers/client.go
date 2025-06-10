package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"insurance-backend/models"
)

type ClientHandler struct {
	DB *gorm.DB
}

func NewClientHandler(db *gorm.DB) *ClientHandler {
	return &ClientHandler{DB: db}
}

func (h *ClientHandler) GetClients(c *fiber.Ctx) error {
	clientType := c.Query("type") // ?type=individual

	var clients []models.Client

	query := h.DB
	if clientType != "" {
		query = query.Where("type = ?", clientType)
	}

	if err := query.Find(&clients).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch clients"})
	}

	return c.JSON(clients)
}

func (h *ClientHandler) GetClient(c *fiber.Ctx) error {
	id := c.Params("id")
	var client models.Client

	if err := h.DB.First(&client, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Client not found"})
	}

	return c.JSON(client)
}

func (h *ClientHandler) CreateClient(c *fiber.Ctx) error {
	var input models.Client

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Проверка на тип
	if input.Type != models.Individual && input.Type != models.Legal {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unknown client type"})
	}

	if err := h.DB.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create client"})
	}

	return c.Status(fiber.StatusCreated).JSON(input)
}

func (h *ClientHandler) UpdateClient(c *fiber.Ctx) error {
	id := c.Params("id")
	var client models.Client
	if err := h.DB.First(&client, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Client not found"})
	}
	if err := c.BodyParser(&client); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	h.DB.Save(&client)
	return c.JSON(client)
}

func (h *ClientHandler) DeleteClient(c *fiber.Ctx) error {
	id := c.Params("id")
	h.DB.Delete(&models.Client{}, id)
	return c.SendStatus(204)
}

package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name      string `json:"name"`
	IsCompany bool   `json:"is_company"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}

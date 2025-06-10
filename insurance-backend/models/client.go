package models

import (
	"gorm.io/gorm"
	"time"
)

type ClientType string

const (
	Individual ClientType = "individual"
	Legal      ClientType = "legal"
)

type Client struct {
	gorm.Model
	Type ClientType `gorm:"type:varchar(20)"`
	// Общие поля
	Address string
	Phone   string

	// Физлицо
	LastName          string // Фамилия
	FirstName         string // Имя
	MiddleName        string // Отчество
	BirthDate         *time.Time
	Gender            string
	DrivingExperience int
	PhotoURL          string

	// Юрлицо
	CompanyName      string
	TIN              string // УНН
	DirectorFullName string
	AccountantName   string

	// Связь с полисами
	Policies []Policy `gorm:"foreignKey:ClientID"`
}

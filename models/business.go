package models

import (
	"time"

	"github.com/google/uuid"
)

type Business struct {
	ID           *uuid.UUID  `gorm:"column:id;primary_key;default:(uuid())" json:"id,omitempty"`
	Alias        string      `gorm:"type:varchar(255);index:idx_business1,unique" json:"alias"`
	Name         string      `gorm:"type:varchar(255)" json:"name"`
	ImageUrl     string      `json:"image_url"`
	IsClosed     bool        `json:"is_closed"`
	Url          string      `json:"url"`
	Coordinates  Coordinates `gorm:"type:string;serializer:json" json:"coordinates"`
	Transactions []string    `gorm:"type:string;serializer:json" json:"transactions"`
	Categories   string      `json:"categories"`
	Price        string      `json:"price"`
	Phone        string      `json:"phone"`
	DisplayPhone string      `json:"display_phone"`
	CreatedAt    *time.Time  `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt    *time.Time  `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
	DeletedAt    *time.Time  `gorm:"type:DATETIME;default:NULL" json:"deleted_at,omitempty"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CreateBusinessBody struct {
	Business
	Location Location `json:"location"`
}

type BusinessSearch struct {
	Alias        string           `gorm:"type:varchar(255);index:idx_business1,unique" json:"alias"`
	Name         string           `json:"name"`
	ImageUrl     string           `json:"image_url"`
	IsClosed     bool             `json:"is_closed"`
	Url          string           `json:"url"`
	Coordinates  Coordinates      `gorm:"type:string;serializer:json" json:"coordinates"`
	Transactions []string         `gorm:"type:string;serializer:json" json:"transactions"`
	Categories   []CategorySearch `gorm:"type:string;serializer:json" json:"categories"`
	Price        string           `json:"price"`
	Phone        string           `json:"phone"`
	DisplayPhone string           `json:"display_phone"`
	Location     LocationSearch   `gorm:"type:string;serializer:json" json:"location"`
	Distance     float64          `json:"distance"`
}

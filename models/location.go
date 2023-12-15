package models

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	ID             *uuid.UUID `gorm:"column:id;primary_key;default:(uuid())" json:"id,omitempty"`
	BusinessID     *uuid.UUID `gorm:"column:business_id;index:idx_loc1;" json:"business_id,omitempty"`
	Business       *Business  `gorm:"foreignKey:business_id" json:"business,omitempty"`
	Address1       string     `json:"address1"`
	Address2       string     `json:"address2"`
	Address3       string     `json:"address3"`
	City           string     `json:"city"`
	ZipCode        string     `json:"zip_code"`
	Country        string     `json:"country"`
	State          string     `json:"state"`
	DisplayAddress []string   `gorm:"type:string;serializer:json" json:"display_address"`
	CreatedAt      *time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt      *time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
	DeletedAt      *time.Time `gorm:"type:DATETIME;default:NULL" json:"deleted_at,omitempty"`
}

type LocationSearch struct {
	Address1       string `json:"address1"`
	Address2       string `json:"address2"`
	Address3       string `json:"address3"`
	City           string `json:"city"`
	ZipCode        string `json:"zip_code"`
	Country        string `json:"country"`
	State          string `json:"state"`
	DisplayAddress any    `json:"display_address"`
}

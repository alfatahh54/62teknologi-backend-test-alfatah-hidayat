package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        *uuid.UUID `gorm:"column:id;primary_key;default:(uuid())" json:"id,omitempty"`
	Alias     string     `gorm:"type:varchar(255);index:idx_cat1,unique" json:"alias"`
	Name      string     `gorm:"type:varchar(255)" json:"name"`
	CreatedAt *time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"type:DATETIME;default:NULL"`
}

type CategorySearch struct {
	Alias string `json:"alias"`
	Name  string `json:"name"`
}

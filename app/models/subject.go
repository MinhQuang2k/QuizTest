package models

import (
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
}

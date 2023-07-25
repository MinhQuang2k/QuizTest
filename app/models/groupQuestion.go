package models

import (
	"gorm.io/gorm"
)

type GroupQuestion struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
	User   *User
}

package models

import (
	"gorm.io/gorm"

	"quiztest/pkg/utils"
)

type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleCustomer UserRole = "customer"
)

type User struct {
	gorm.Model
	FullName string   `json:"full_name"`
	Email    string   `json:"email" gorm:"unique;not null;index:idx_user_email"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Password = utils.HashAndSalt([]byte(user.Password))
	if user.Role == "" {
		user.Role = UserRoleCustomer
	}
	return nil
}

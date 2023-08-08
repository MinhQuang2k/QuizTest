package interfaces

import (
	"gorm.io/gorm"
)

// IDatabase interface
type IDatabase interface {
	GetInstance() *gorm.DB
}

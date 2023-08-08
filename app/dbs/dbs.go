package dbs

import (
	"quiztest/app/interfaces"
	"quiztest/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"quiztest/config"
)

type database struct {
	db *gorm.DB
}

func NewDatabase() interfaces.IDatabase {
	cfg := config.GetConfig()
	db, err := gorm.Open(mysql.Open(cfg.DatabaseURI), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	// Set up connection pool
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)

	return &database{
		db: db,
	}
}

// GetInstance get database instance
func (d *database) GetInstance() *gorm.DB {
	return d.db
}

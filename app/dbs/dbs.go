package dbs

import (
	"quiztest/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"quiztest/app/models"
	"quiztest/config"
)

var Database *gorm.DB

func Init() {
	cfg := config.GetConfig()
	database, err := gorm.Open(mysql.Open(cfg.DatabaseURI), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	// Set up connection pool
	sqlDB, err := database.DB()
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)
	Database = database

	Migrate()
}

func Migrate() {
	User := models.User{}
	Category := models.Category{}
	Exam := models.Exam{}
	GroupQuestion := models.GroupQuestion{}
	Question := models.Question{}
	Subject := models.Subject{}
	Database.AutoMigrate(&User, &GroupQuestion, &Category, &Subject, &Exam, &Question)
}

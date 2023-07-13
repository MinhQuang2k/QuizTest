package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	dbUser := "root"
	dbPass := "root"
	dbName := "quiztest"
	dbHost := "localhost"
	dbPort := "3306"

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connect database")
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

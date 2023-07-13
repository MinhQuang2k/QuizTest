package models

import (
	"blog.com/common"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&PostModel{})
	db.AutoMigrate(&CommentModel{})
	db.AutoMigrate(&GroupQuestionModel{})
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

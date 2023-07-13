package models

import (
	"blog.com/common"
	"gorm.io/gorm"
)

type GroupQuestionModel struct {
	gorm.Model
	Name   string `gorm:"unique;size:2048"`
	User   UserModel
	UserID uint
}

func (model *GroupQuestionModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Updates(data).Error
	return err
}

func DeleteGroupQuestion(id int) error {
	db := common.GetDB()
	var models GroupQuestionModel
	err := db.First(&models, id).Error

	if err == nil {
		db.Delete(&models)
	}

	return err
}

func FindGroupQuestion(id int) (GroupQuestionModel, error) {
	db := common.GetDB()
	var models GroupQuestionModel
	err := db.First(&models, id).Error

	return models, err
}

func FindGroupQuestionPaging(offset, size int, userID uint) ([]GroupQuestionModel, int64, error) {
	db := common.GetDB()
	var models []GroupQuestionModel
	var count int64
	tx := db.Begin()

	tx.Model(GroupQuestionModel{}).Where("user_id = ?", userID).Count(&count)
	tx.Where(GroupQuestionModel{UserID: userID}).Preload("GroupQuestions").Offset(offset).Limit(size).Find(&models)

	err := tx.Commit().Error
	return models, count, err
}

func FindGroupQuestionAll(userID uint) ([]GroupQuestionModel, error) {
	db := common.GetDB()
	var models []GroupQuestionModel

	tx := db.Begin()
	tx.Where(GroupQuestionModel{UserID: userID}).Preload("GroupQuestions").Find(&models)

	err := tx.Commit().Error
	return models, err
}

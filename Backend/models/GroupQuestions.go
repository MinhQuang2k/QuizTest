package models

import (
	"strconv"

	"blog.com/common"
	"gorm.io/gorm"
)

type GroupQuestionModel struct {
	gorm.Model
	Name   string `gorm:"unique,size:2048"`
	User   UserModel
	UserID uint
}

func (model *GroupQuestionModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Updates(data).Error
	return err
}

func DeleteGroupQuestion(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(&GroupQuestionModel{}).Error
	return err
}

func FindGroupQuestionAll(condition interface{}) (GroupQuestionModel, error) {
	db := common.GetDB()
	var model GroupQuestionModel
	tx := db.Begin()
	tx.Where(condition).First(&model)
	tx.Model(&model).Association("User").Find(&model.User)
	err := tx.Commit().Error
	return model, err
}

func FindGroupQuestionPaging(all, limit, offset string, userID uint) ([]GroupQuestionModel, int64, error) {
	db := common.GetDB()
	var models []GroupQuestionModel
	var count int64

	offset_int, err := strconv.Atoi(offset)
	if err != nil {
		offset_int = 0
	}

	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		limit_int = 20
	}

	tx := db.Begin()
	count = tx.Where(GroupQuestionModel{UserID: userID}).Association("GroupQuestions").Count()
	if all != "" {
		tx.Where(GroupQuestionModel{UserID: userID}).Preload("GroupQuestions").Offset(offset_int).Limit(limit_int).Find(&models)
	} else {
		tx.Where(GroupQuestionModel{UserID: userID}).Preload("GroupQuestions").Find(&models)
	}

	err = tx.Commit().Error
	return models, count, err
}

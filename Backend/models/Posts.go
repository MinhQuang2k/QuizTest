package models

import (
	"strconv"

	"blog.com/common"
	"gorm.io/gorm"
)

type PostModel struct {
	gorm.Model
	Slug        string `gorm:"unique_index"`
	Title       string
	Description string `gorm:"size:2048"`
	Body        string `gorm:"size:2048"`
	User        UserModel
	UserID      uint
	Comments    []CommentModel `gorm:"ForeignKey:PostID"`
}

type CommentModel struct {
	gorm.Model
	Post   PostModel
	PostID uint
	User   UserModel
	UserID uint
	Body   string `gorm:"size:2048"`
}

func (model *PostModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Updates(data).Error
	return err
}

func (self *PostModel) GetComments() error {
	db := common.GetDB()
	tx := db.Begin()
	tx.Where(PostModel{UserID: self.UserID}).Preload("Comments").Find(&self)
	err := tx.Commit().Error
	return err
}

func DeletePost(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(&PostModel{}).Error
	return err
}

func DeleteComment(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(&CommentModel{}).Error
	return err
}

func FindPost(condition interface{}) (PostModel, error) {
	db := common.GetDB()
	var model PostModel
	tx := db.Begin()
	tx.Where(condition).First(&model)
	tx.Model(&model).Association("User").Find(&model.User)
	err := tx.Commit().Error
	return model, err
}

func FindPostPaging(author, limit, offset string) ([]PostModel, int64, error) {
	db := common.GetDB()
	var models []PostModel
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
	if author != "" {
		var userModel UserModel
		tx.Where(UserModel{Username: author}).First(&userModel)

		if userModel.ID != 0 {
			count = tx.Model(&userModel).Association("Posts").Count()
			tx.Where(PostModel{UserID: userModel.ID}).Preload("Posts").Offset(offset_int).Limit(limit_int).Find(&models)
		}
	} else {
		db.Model(&models).Count(&count)
		db.Offset(offset_int).Limit(limit_int).Find(&models)
	}

	err = tx.Commit().Error
	return models, count, err
}

package models

import (
	"errors"

	"blog.com/common"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID           uint        `gorm:"primary_key"`
	Username     string      `gorm:"column:username"`
	Email        string      `gorm:"column:email;unique_index"`
	Bio          string      `gorm:"column:bio;size:1024"`
	Image        *string     `gorm:"column:image"`
	PasswordHash string      `gorm:"column:password;not null"`
	Posts        []PostModel `gorm:"ForeignKey:UserID"`
}

func (u *UserModel) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *UserModel) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (model *UserModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindUser(condition interface{}) (UserModel, error) {
	db := common.GetDB()
	var model UserModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

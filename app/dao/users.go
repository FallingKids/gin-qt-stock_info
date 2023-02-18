package dao

import (
	"time"

	"github.com/gin-qt-business/app/config"
)

const TABLE_NAME = "Users"

type UserDao struct {
	ID        uint `gorm:"primary_key"`
	Uid       string
	Username  string
	Password  string
	Phone     string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt uint
}

func (user *UserDao) AddUser() error {
	err := config.DB.Table(TABLE_NAME).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *UserDao) GetUser() error {
	err := config.DB.Table(TABLE_NAME).Take(&user).Error
	if err != nil {
		return err
	}
	return nil
}

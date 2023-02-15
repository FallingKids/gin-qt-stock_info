package dao

import (
	"time"

	"github.com/gin-qt-business/app/config"
)

const TABLE_NAME = "users"

type User struct {
	ID        uint `gorm:"primary_key"`
	Uid       string
	Username  string
	Password  string
	Phone     string
	CreatedAt time.Time `gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano"`
}

func AddUser(user User) (res *User, err error) {
	err = config.DB.Table(TABLE_NAME).Create(user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

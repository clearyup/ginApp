package models

import (
	"ginProject/dao"
)

type User struct {
	ID        int    `gorm:"column:id;" `
	UserName  string `gorm:"column:username;index:i_user;" `
	Password  string `gorm:"column:password;" `
	Email     string `gorm:"column:email;" `
	CreatedAt string `gorm:"column:created_at;" `
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) GetUserByName(name string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", name).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

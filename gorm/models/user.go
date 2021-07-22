package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string
	UserAddress []*UserAddress `gorm:"ForeignKey:user_id"`
}

func (User) TableName() string {
	return "user"
}

type UserAddress struct {
	gorm.Model
	UserId   int64
	Contacts string
	Tel      string
	Address  string
}

func (UserAddress) TableName() string {
	return "user_address"
}

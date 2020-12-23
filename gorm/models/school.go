package models

import (
	"github.com/jinzhu/gorm"
)

type Teacher struct {
	gorm.Model
	Name    string
	Student []*Student `gorm:"ForeignKey:teacher_id"`
}

func (Teacher) TableName() string {
	return "s_teacher"
}

type Student struct {
	gorm.Model
	TeacherId int64
	Name      string
}

func (Student) TableName() string {
	return "s_student"
}

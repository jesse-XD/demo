package gorm

import (
	"demo/gorm/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type School struct {
}

func (s *School) Add(t *models.Teacher) (errs []error) {
	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	if errs = db.Create(t).GetErrors(); len(errs) > 0{
		return
	}
	return
}

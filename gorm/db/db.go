package db

import (
	"github.com/jinzhu/gorm"
	"sync"
)


var (
	db *gorm.DB
	once sync.Once
)

// 连接数据库
func DB() *gorm.DB{
	once.Do(func() {
		d, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}
		d.LogMode(true)
		db = d
	})
	return db
}

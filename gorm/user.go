package gorm

import (
	"demo/gorm/db"
	"demo/gorm/models"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
}

// 添加
func (u *User) Add(user *models.User) (errs []error) {
	if errs = db.DB().Create(user).GetErrors(); len(errs) > 0 {
		return
	}
	return
}

// GetById 根据 id 获取数据
func (u *User) GetById(id int64, user *models.User) bool {
	if db.DB().First(&user, id).RecordNotFound() {
		return false
	}

	return true
}

// Update 修改数据
func (u *User) Update(user *models.User) (errs []error) {
	// 先判断是否存在
	var old models.User
	if !u.GetById(int64(user.ID), &old) {
		errs = []error{errors.New("记录不存在")}
		return
	}
	// 修改
	if errs = db.DB().Model(&old).Save(user).GetErrors(); len(errs) > 0 {
		return
	}

	return
}

// Delete 删除 force：false|true 是否物理删除
func (u *User) Delete(id int64, force bool) (errs []error) {
	// 先判断是否存在
	var old models.User
	if !u.GetById(id, &old) {
		errs = []error{errors.New("记录不存在")}
		return
	}
	d := db.DB()
	if force {
		// 物理删除
		d = d.Unscoped()
	}
	if errs = d.Delete(&old).GetErrors(); len(errs) > 0 {
		return
	}

	return
}

// List 列表
func (u *User) List(param map[string]string) (userList []*models.User) {
	d := db.DB()
	if val, ok := param["name"]; ok && val != "" {
		d = d.Where("name like ?", "%" + val + "%")
	}
	d = d.Preload("UserAddress", func(db *gorm.DB)*gorm.DB {
		if val, ok := param["contacts"]; ok && val != "" {
			db = db.Where("contacts like ?", "%" + val + "%")
		}
		if val, ok := param["tel"]; ok && val != "" {
			db = db.Where("tel = ?", val)
		}
		return db
	}).Find(&userList)
	return
}

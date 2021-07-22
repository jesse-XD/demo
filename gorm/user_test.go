package gorm

import (
	"demo/gorm/models"
	"encoding/json"
	"testing"
)

func TestUser_Add(t *testing.T) {
	var user models.User
	user.Name = "Tom"
	var addressList []*models.UserAddress
	addressList = append(addressList, &models.UserAddress{Contacts: "Jesse", Tel: "15123123121", Address: "广州123"})
	addressList = append(addressList, &models.UserAddress{Contacts: "Jesse2", Tel: "18123123121", Address: "深圳123"})
	user.UserAddress = addressList
	errs := (*User)(nil).Add(&user)
	if len(errs) > 0 {
		t.Fatal(errs)
	}
}

func TestUser_Update(t *testing.T) {
	var user models.User
	user.ID = 2
	user.Name = "Jack"
	//var addressList []*models.UserAddress
	//addressList = append(addressList, &models.UserAddress{Model: gorm.Model{ID: 1},Contacts: "Jesse11", Tel: "15123123121", Address: "广州123"})
	//addressList = append(addressList, &models.UserAddress{Contacts: "Jesse2", Tel: "18123123121", Address: "深圳123"})
	//user.UserAddress = addressList
	errs := (*User)(nil).Update(&user)
	if len(errs) > 0 {
		t.Fatal(errs)
	}
}

func TestUser_Delete(t *testing.T) {
	id := 3
	errs := (*User)(nil).Delete(int64(id), true)
	if len(errs) > 0 {
		t.Fatal(errs)
	}
}

func TestUser_List(t *testing.T) {
	param := make(map[string]string)
	param["name"] = "to"
	param["contacts"] = "Jesse2"
	list := (*User)(nil).List(param)
	s, _ := json.Marshal(list)
	t.Log(string(s))
}
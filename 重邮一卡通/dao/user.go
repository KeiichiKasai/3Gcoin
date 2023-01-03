package dao

import (
	"fmt"
	"main.go/model"
)

// SearchUserByUsername 通过用户名来查找用户
func SearchUserByUsername(username string) (u model.User, err error) {
	tx := DB.Where("username=?", username).Find(&u)
	err = tx.Error
	if err != nil {
		fmt.Printf("err:%v", err)
		panic(err)
	}
	return u, err
}

// InsertUser 插入用户数据
func InsertUser(u model.User) (err error) {
	tx := DB.Create(&u)
	err = tx.Error
	if err != nil {
		fmt.Printf("err:%v", err)
		panic(err)
	}
	return err
}

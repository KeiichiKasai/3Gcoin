package dao

import (
	"fmt"
	"main.go/model"
)

// GetMoney 获取用户余额
func GetMoney(username string) (money int) {
	u, err := SearchUserByUsername(username)
	if err != nil {
		fmt.Printf("err:%v", err)
		panic(err)
	}
	money = u.Money
	return money
}

// UpdateMoney 更新用户余额
func UpdateMoney(username string, money int) (err error) {
	var u model.User
	tx := DB.Model(&u).Where("username=?", username).Update("money", money)
	err = tx.Error
	if err != nil {
		fmt.Printf("err:%v", err)
		panic(err)
	}
	return err
}

// InsertRecord 插入交易记录
func InsertRecord(r model.Record) (err error) {
	tx := DB.Create(&r)
	err = tx.Error
	return err
}

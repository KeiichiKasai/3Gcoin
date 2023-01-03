package service

import (
	"fmt"
	"main.go/dao"
	"main.go/model"
)

func AddMoney(username string, money int) (err error) {
	before := dao.GetMoney(username)
	now := before + money
	err = dao.UpdateMoney(username, now)
	return err
}
func MinusMoney(username string, money int) (flag bool, err error) {
	before := dao.GetMoney(username)
	flag = false
	now := before - money
	if now < 0 {
		flag = true
		return flag, nil
	}

	err = dao.UpdateMoney(username, now)
	return flag, err

}
func CreateRecord(r model.Record) (err error) {
	err = dao.InsertRecord(r)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	return err
}

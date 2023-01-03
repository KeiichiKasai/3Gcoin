package service

import (
	"main.go/dao"
	"main.go/model"
)

// CheckUser 用来检测是否重名,重名返回true，否则返回false
func CheckUser(username string) (flag bool) {
	u, _ := dao.SearchUserByUsername(username)
	flag = true
	if u.Username != username {
		flag = false
		return flag
	}
	return flag
}

// CreateUser 创建新用户
func CreateUser(u model.User) (err error) {
	err = dao.InsertUser(u)
	return err
}

func CheckPassword(username string, password string) (flag bool) {
	u, _ := dao.SearchUserByUsername(username)
	flag = true
	if password != u.Password {
		flag = false
		return flag
	}
	return flag
}
func GetUId(username string) (UId int) {
	u, _ := dao.SearchUserByUsername(username)
	return u.Id
}

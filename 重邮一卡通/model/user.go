package model

type User struct { //用户结构体
	Id       int    `json:"id"`       //用户id
	Username string `json:"username"` //用户名
	Password string `json:"password"` //密码
	Money    int    `json:"money"`    //余额
}

package model

import "time"

type Record struct {
	Rid       int       `json:"rid"`       //交易id
	Uid       int       `json:"uid"`       //转账者id
	CreatedAt time.Time `json:"CreatedAt"` //转账时间
	Context   string    `json:"context"`   //备注
	Money     int       `json:"money"`     //转账金额
	Object    string    `json:"object"`    //转账对象
}

package service

import (
	"main.go/dao"
	"main.go/model"
)

// GetRecordByRid 通过Rid获取Record
func GetRecordByRid(rid int) (r model.Record, err error) {
	r, err = dao.SearchRecordByRid(rid)
	return r, err
}

// AttachContext 通过Uid修改备注
func AttachContext(rid int, uid int, context string) (err error) {
	err = dao.InsertContext(rid, uid, context)
	return err
}

// GetRecordByContext 通过Context查找Record
func GetRecordByContext(uid int, context string) (r []model.Record, err error) {
	r, err = dao.SearchRecordByContext(uid, context)
	return r, err
}

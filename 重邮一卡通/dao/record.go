package dao

import "main.go/model"

func SearchRecordByRid(rid int) (r model.Record, err error) {
	tx := DB.Where("rid=?", rid).Find(&r)
	err = tx.Error
	return r, err
}

func InsertContext(rid int, uid int, context string) (err error) {
	var r model.Record
	tx := DB.Model(&r).Where("rid=?", rid).Where("uid=?", uid).Update("context", context)
	err = tx.Error
	return err
}

func SearchRecordByContext(uid int, context string) (r []model.Record, err error) {
	tx := DB.Where("uid=?", uid).Where("context LIKE ?", context).Find(&r)
	err = tx.Error
	return r, err
}

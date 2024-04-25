package dao

import (
	"AHUT-Forum/config"
	logger "AHUT-Forum/config/log"
	"reflect"
)

var mysql = config.Mysql{}
var db = mysql.Db

//Insert
func Insert(tablename string, model interface{}) {
	mysql.Mu.RLock()
	defer mysql.Mu.RUnlock()

	if tablename == "" {
		return
	}

	if isEmptyModel(model) {
		return
	}

	err := db.Table(tablename).Create(model).Error
	logger.Assert(err)
}

//Update
func Update(tablename string, model interface{}) {
	if tablename == "" {
		return
	}

	if isEmptyModel(model) {
		return
	}

	err := db.Table(tablename).Save(model).Error
	logger.Assert(err)
}

//Select
func Select(tablename string, id uint, model interface{}) []interface{} {
	mysql.Mu.RLock()
	defer mysql.Mu.RUnlock()
	if tablename == "" {
		return nil
	}

	if isEmptyModel(model) {
		return nil
	}

	var records []interface{}
	err := db.Table(tablename).First(model, id).Error
	logger.Assert(err)
	return records
}

//Delete
func Delete(tablename string, model interface{}, id uint) {
	if tablename == "" {
		return
	}

	if isEmptyModel(model) {
		return
	}

	err := db.Table(tablename).Delete(model, id).Error
	logger.Assert(err)
}

// 判断结构体是否为空
func isEmptyModel(model interface{}) bool {
	// 使用反射获取结构体的值
	val := reflect.ValueOf(model)

	// 如果结构体是指针类型，则获取指针指向的值
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// 遍历结构体的字段
	for i := 0; i < val.NumField(); i++ {
		// 获取字段的值
		fieldValue := val.Field(i)

		// 使用反射判断字段的值是否为零值
		if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
			return true
		}
	}

	// 如果所有字段都不为空，则返回false
	return false
}

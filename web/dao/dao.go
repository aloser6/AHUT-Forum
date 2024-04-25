package dao

import (
	"AHUT-Forum/config"
	logger "AHUT-Forum/config/log"
)

var mysql = config.Mysql{}
var db = mysql.Db

//Insert
func Insert(tablename string, model interface{}) {
	mysql.Mu.Lock()
	defer mysql.Mu.Unlock()

	if tablename == "" || model == nil {
		return
	}

	err := db.Table(tablename).Create(model).Error
	logger.Assert(err)
}

//Update
func Update(tablename string, model interface{}) {
	mysql.Mu.Lock()
	defer mysql.Mu.Unlock()

	if tablename == "" || model == nil {
		return
	}

	err := db.Table(tablename).Save(model).Error
	logger.Assert(err)
}

//Select
func Select(tablename string, id uint, model interface{}) []interface{} {
	mysql.Mu.RLock()
	defer mysql.Mu.RUnlock()
	if tablename == "" || model == nil {
		return nil
	}

	var records []interface{}
	err := db.Table(tablename).First(model, id).Error
	logger.Assert(err)
	return records
}

//Delete
func Delete(tablename string, model interface{}, id uint) {
	mysql.Mu.Lock()
	defer mysql.Mu.Unlock()

	if tablename == "" || model == nil {
		return
	}

	err := db.Table(tablename).Delete(model, id).Error
	logger.Assert(err)
}

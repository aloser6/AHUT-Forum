package dao

import (
	logger "AHUT-Forum/config/log"
)

var db Mysql

//Insert
func Insert(tablename string, model interface{}) {
	if tablename == "" || model == nil {
		return
	}

	db.Mu.Lock()
	defer db.Mu.Unlock()

	err := db.Db.Table(tablename).Create(model).Error
	logger.Assert(err)
}

//Update
func Update(tablename string, model interface{}) {
	if tablename == "" || model == nil {
		return
	}

	db.Mu.Lock()
	defer db.Mu.Unlock()

	err := db.Db.Table(tablename).Save(model).Error
	logger.Assert(err)
}

//Select
func Select(tablename string, id uint, model interface{}) []interface{} {
	if tablename == "" || model == nil {
		return nil
	}

	db.Mu.RLock()
	defer db.Mu.RUnlock()

	var records []interface{}
	err := db.Db.Table(tablename).First(model, id).Error
	logger.Assert(err)
	return records
}

//Delete
func Delete(tablename string, model interface{}, id uint) {
	if tablename == "" || model == nil {
		return
	}

	db.Mu.Lock()
	defer db.Mu.Unlock()

	err := db.Db.Table(tablename).Delete(model, id).Error
	logger.Assert(err)
}

package dao

import (
	"AHUT-Forum/config"
	logger "AHUT-Forum/config/log"
)

var mysql = config.Mysql{}
var db = mysql.Db

//Insert
func Insert(tablename string, model interface{}) error {
	err := db.Table(tablename).Create(model).Error
	if err != nil {
		logger.Assert(err)
	}
	return err
}

func InsertLock(tablename string, model interface{}) {
	mysql.Mu.RLock()
	defer mysql.Mu.Unlock()
	Insert(tablename, model)
}

//Update
func Updata(tablename string, model interface{}) error {
	err := db.Table(tablename).Save(model).Error
	if err != nil {
		logger.Assert(err)
	}
	return err
}

//Select
func Select(tableName string, id uint, model interface{}) error {
	err := db.Table(tableName).First(model, id).Error
	if err != nil {
		logger.Assert(err)
	}
	return err
}

func SelectLock(tableName string, id uint, model interface{}) {
	mysql.Mu.RLock()
	defer mysql.Mu.Unlock()
	Select(tableName, id, model)
}

//Delete
func Delete(tableName string, model interface{}, id uint) error {
	err := db.Table(tableName).Delete(model, id).Error
	if err != nil {
		logger.Assert(err)
	}
	return err
}

package dao

import (
	"AHUT-Forum/config"
	logger "AHUT-Forum/config/log"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Db *gorm.DB
	Mu sync.RWMutex
}

func Mysql_init(conf config.Config) {
	db = Mysql{}
	var err error
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Ip, conf.Mysql.Port, conf.Mysql.Dbname)
	logger.Info(url)
	if db.Db == nil {
		db.Db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	}
	logger.Assert(err)
}

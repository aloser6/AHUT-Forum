package config

import (
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

func (mq *Mysql) Mysql_init(conf *Config) {
	var err error
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Ip, conf.Mysql.Port, conf.Mysql.Dbname)
	logger.Info(url)
	mq.Db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	logger.Assert(err)
}

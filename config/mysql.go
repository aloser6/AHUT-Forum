package config

import (
	logger "AHUT-Forum/config/log"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Db *gorm.DB
}

func (mq *Mysql) Mysql_init() {
	var err error
	config := Config{}
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.Mysql.Username, config.Mysql.Password, config.Mysql.Ip, config.Mysql.Port, config.Mysql.Dbname)
	mq.Db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	logger.Assert(err)
}

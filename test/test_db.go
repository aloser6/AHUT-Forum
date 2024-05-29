package main

import (
	"AHUT-Forum/config"
	"AHUT-Forum/units"
	"AHUT-Forum/web/dao"
)

func test_db() {
	units.Init()
	conf := config.Config{}
	conf.Config_init()
	dao.Mysql_init(conf)
}

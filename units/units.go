package units

import (
	"AHUT-Forum/config"
	logger "AHUT-Forum/config/log"
	"AHUT-Forum/web/dao"
)

func Init() {
	logger.Logger_init()
	conf := config.Config{}
	conf.Config_init()
	dao.Mysql_init(conf)
}

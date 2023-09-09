package mysql

import (
	"ISPS/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(s string) *gorm.DB {
	y := config.Yaml{}
	DB, err := gorm.Open(mysql.Open(y.ReadYamlString("mysql."+s+".dsn")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

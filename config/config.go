package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Yaml struct {
}

//查询
func (y *Yaml) ReadYaml(Key string) string { //Key 路径
	config := viper.New()
	InitReadconfig(config)
	return config.GetString(Key)
}

/*可以完成增，删，改
如果输入的key不存在则增
如果输入的value为空则删除
如果输入的key存在则修改
*/
func (y *Yaml) SetYaml(Key string, Value string) { //Key 路径 Value值
	config := viper.New()
	InitReadconfig(config)
	config.Set(Key, Value)
	InitSetconfig(config)
}

func (y *Yaml) DeleteYaml(Key string) {
	config := viper.New()
	InitReadconfig(config)
	config.Set(Key, nil)
	InitSetconfig(config)
}

//用于找到文件
func InitReadconfig(config *viper.Viper) {
	config.SetConfigFile("/home/kasadin/go_test/zhangchi/ISPS/config.yaml") //yaml文件的绝对路径
	err := config.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

//用于修改数据后保存
func InitSetconfig(config *viper.Viper) {
	err := config.WriteConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var DB *gorm.DB

//用于连接数据库
func InitMySQL(y Yaml) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(y.ReadYaml("mysql.dsn")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

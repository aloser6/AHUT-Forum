package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Yaml struct {
	Key   string
	Value string
}

//查询
func (y *Yaml) ReadYaml() string {
	config := viper.New()
	InitReadconfig(config)
	return config.GetString(y.Key)
}

//可以完成增，删，改
//如果输入的key不存在则增
//如果输入的value为空则删除
//如果输入的key存在则修改
func (y *Yaml) SetYaml(value string) {
	config := viper.New()
	InitReadconfig(config)
	y.Value = value
	config.Set(y.Key, y.Value)
	InitSetconfig(config)
}

func (y *Yaml) DeleteYaml() {
	config := viper.New()
	InitReadconfig(config)
	config.Set(y.Key, nil)
	InitSetconfig(config)
}

//用于找到文件
func InitReadconfig(config *viper.Viper) {
	config.SetConfigName("app")
	config.AddConfigPath("config")
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

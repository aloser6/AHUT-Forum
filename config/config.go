package config

import (
	logger "AHUT-Forum/config/log"
	"os"

	"gopkg.in/yaml.v2"
)

var config_path = "../config.yml"

type Config struct {
	Mysql mysql_config `yaml:"mysql"` //yml反射不能是私有字段故大写
}

type mysql_config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Ip       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
}

func (conf *Config) Init() {
	logger.Logger_init()
	conf.config_init()
}

func (conf *Config) config_init() {
	data, err := os.ReadFile(config_path)
	logger.Assert(err)
	err = yaml.Unmarshal(data, &conf)
	logger.Info(conf.Mysql.Username)
	logger.Assert(err)
	logger.Info("config → %+v\n", conf)
}

// config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}

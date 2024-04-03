package main

import "AHUT-Forum/config"

func test_db() {
	conf := config.Config{}
	conf.Init()
	_ = config.Mysql{}
}

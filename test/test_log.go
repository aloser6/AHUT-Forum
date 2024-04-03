package main

import (
	logger "AHUT-Forum/config/log"
)

func test_log() {
	logger.Logger_init()
	str := "it's a test"
	logger.Info("logger_info %v", str)
	var err error
	logger.Assert(err)
}

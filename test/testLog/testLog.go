package main

import (
	"ISPS/log"
	"time"
)

func main() {
	log1 := log.NewLog("unknown", "./", "tangbLog.log", 1024*2)
	for {
		log1.Debug("这是一条Debug日志", "consolepath")
		log1.Info("这是一条Info日志", "consolepath")
		log1.Warn("这是一条Warning日志", "filepath")
		log1.Error("这是一条Error日志", "")
		log1.Fatal("这是一条Fatal日志", "")
		time.Sleep(time.Second)
	}
}

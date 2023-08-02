package main

import (
	"ISPS/log"
	"time"
)

func main() {
	/*参1：日志等级
	*参2：文件路径
	*参3：文件名称
	*参4：文件最大容量*/
	//y := config.Yaml{}
	// fmt.Println(y.ReadYaml("log.maxFileSize"))
	// c1, err1 := strconv.Atoi(y.ReadYaml("log.maxFileSize"))
	// if err1 != nil {
	// 	fmt.Println("log error")
	// }
	// c2 := int64(c1)
	// fmt.Print(c2)
	//log1 := log.NewLog(y.ReadYamlString("log.levelstr"), y.ReadYamlString("log.fp"), y.ReadYamlString("log.fn"), y.ReadYamlInt64("log.maxFileSize"))
	// log1.Debug("这是一条Debug日志", "consolepath")
	// log1.Info("这是一条Info日志", "consolepath")
	// log1.Warn("这是一条Warning日志", "filepath")
	// log1.Error("这是一条Error日志", "")
	// log1.Fatal("这是一条Fatal日志", "")
	log1 := log.NewLog()
	for {
		log1.Warn("这是一条Warning日志", "filepath")
		log1.Error("这是一条Error日志", "")
		log1.Fatal("这是一条Fatal日志", "")
		time.Sleep(time.Second)
	}
}

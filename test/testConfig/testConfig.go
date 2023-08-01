package main

import (
	"ISPS/config"
	"fmt"
)

func main() {
	y := config.Yaml{}
	//utils.InitMySQL(y)
	//y1 := utils.Yaml{Key: "mysql.a", Value: "a"}
	//fmt.Println(y.ReadYaml())
	//fmt.Println(y1.ReadYaml())
	//y1.SetYaml()
	fmt.Println(y.ReadYaml("mysql.dsn"))
	//fmt.Println(y1.ReadYaml())
}

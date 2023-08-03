package main

import (
	"ISPS/config"
	"fmt"
)

func main() {
	y := config.Yaml{}
	//config.InitMySQL(y)
	//y1 := utils.Yaml{Key: "mysql.a", Value: "a"}
	//fmt.Println(y.ReadYaml())
	//fmt.Println(y1.ReadYaml())
	//y1.SetYaml()
	fmt.Print(y.ReadYamlString("mysql.dsn"))
}

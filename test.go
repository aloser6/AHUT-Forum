package main

import (
	"fmt"
	"main/go_test/zhangchi/ISPS/utils"
)

func main() {
	//y := utils.Yaml{Key: "mysql.dsn"}
	y1 := utils.Yaml{Key: "mysql.a", Value: "a"}
	//fmt.Println(y.ReadYaml())
	fmt.Println(y1.ReadYaml())
	y1.SetYaml()
	//fmt.Println(y.ReadYaml())
	fmt.Println(y1.ReadYaml())
}

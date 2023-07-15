package main

import (
	"fmt"
	"main/go_test/zhangchi/ISPS/utils"
)

func main() {
	y := utils.Yaml{Key: "mysql.dsn"}
	fmt.Println(y.ReadYaml())
	y.SetYaml("aaa")
	fmt.Println(y.ReadYaml())
}

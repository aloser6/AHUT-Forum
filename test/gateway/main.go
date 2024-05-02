package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	Init_route(r)
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

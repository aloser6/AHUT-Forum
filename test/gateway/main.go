package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	Init_route(r)
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init_route(r *gin.Engine) {
	r.POST("/ahutforum/user/account/login", test_POST)
}

var body struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func test_POST(c *gin.Context) {
	//获取请求体的json数据
	if err := c.BindJSON(&body); err != nil {
		return
	}
	account := body.Account
	password := body.Password
	fmt.Println()
	fmt.Println(account)
	fmt.Println(password)
	fmt.Println()

	//返回
	c.JSON(200, gin.H{
		"status_code":    200,
		"status_message": "Success",
		"token":          "xxxx",
	})
}

//获取 url 的参数https://www.w3cschool.cn/golang_gin/golang_gin-6dpf3ls1.html

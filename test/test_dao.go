package main

import (
	"AHUT-Forum/models"
	"AHUT-Forum/web/dao"
	"fmt"
)

func test_dao() {
	newUser := models.User{
		Account:  12345,
		Password: "password123",
		Username: "wangwu",
		Sex:      "男",
		Grade:    "大一",
		College:  "计算机学院",
		Major:    "软件工程",
	}

	//增
	dao.Insert(&newUser)
	fmt.Println("用户数据增加成功")

	//改
	updatedUser := models.User{
		Account:  10086,
		Password: "123password",
		Username: "lixiang",
		Sex:      "女",
		Grade:    "大二",
		College:  "材料科学与工程学院",
		Major:    "无机非金属",
	}
	dao.Updata(&updatedUser)
	fmt.Println("用户数据修改成功")

	//查
	id := uint(1)
	findUser := models.User{}
	resUser, err := dao.Select(id, &findUser)
	if err != nil {
		fmt.Printf("查询用户失败：%v\n", err)
	} else {
		fmt.Printf("查询到的用户信息：%+v\n", resUser)
	}

	//删
	uid := uint(1)
	deletedUser := models.User{}
	dao.Delete(uid, &deletedUser)
	fmt.Println("用户数据删除成功")

}

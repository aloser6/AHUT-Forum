package main

import (
	"AHUT-Forum/units"
	"AHUT-Forum/web/dao"
	"AHUT-Forum/web/models"
	"fmt"
	"time"
)

func test_dao() {
	units.Init()

	// Test Insert
	fmt.Println("测试插入...")
	user := models.User{
		Account:   123456,
		Password:  "password123",
		Username:  "Test User",
		StartTime: time.Now(),
		Sex:       models.Male,
		Grade:     "21",
		College:   "Test College",
		Major:     "Test Major",
	}
	dao.Insert("users", &user)
	fmt.Println("插入成功!")

	// Test Update
	fmt.Println("测试更新...")
	user.Username = "Updated Test User"
	dao.Update("users", &user)
	fmt.Println("更新成功!")

	// Test Select
	fmt.Println("测试查询...")
	selectedUser := models.User{}
	records := dao.Select("users", uint(user.UID), &selectedUser)
	if len(records) > 0 {
		fmt.Println("Selected User:", selectedUser)
	} else {
		fmt.Println("未找到用户。")
	}
	fmt.Println("查询成功!")

	// Test Delete
	// fmt.Println("测试删除...")
	// dao.Delete("users", &user, uint(user.UID))
	// fmt.Println("删除成功!")
}

package main

import (
	"ISPS/dao/mysql"
	"fmt"
)

func main() {
	/*testUser*/
	// y := config.Yaml{}
	// db := mysql.InitUser(y)
	// var account mysql.Accounts
	// account = mysql.Accounts{
	// 	Account:  "1",
	// 	Password: "123",
	// 	Email:    "2",
	// }
	// err := db.Where("account_id = ?", 100).Find(&account).Error
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err1 := db.Create(&account).Error
	// if err1 != nil {
	// 	fmt.Println(err1)
	// }
	// fmt.Println(account.AccountId)
	// u := mysql.InitUser()
	// var cp mysql.ConcernPlates
	// u.First(&cp)
	// //u.Where("account_id", 1).Find(&cp)
	// fmt.Println(cp.ConcernPlatesID)
	//147129
	//10010011 10000001 00000100
	var user mysql.UserDB
	// if user.FindConcernPlate(1, 1) {
	// 	fmt.Println(1)
	// } else {
	// 	fmt.Println(2)
	// }
	var cp mysql.ConcernPlates
	cp.AccountId = 2
	//fmt.Println(len("1471290040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"))
	//flag := user.InsertConcernPlate(&cp, 298)
	// if user.DeleteConcernPlates(&cp, 298) {
	// 	fmt.Println(cp.ConcernPlatesID)
	// 	fmt.Println(len(cp.ConcernPlatesID))
	// } else {
	// 	fmt.Println(1)
	// }
	if user.SelectConcernPlate(&cp, 1) {
		fmt.Println(1)
	}
	//user.InsertConcernPlate(&cp, 19)
	//10010011 10000001 00000111
	accounts := user.SelectAllConcernPlates(&cp)
	for i := 0; i < cp.ConcernPlatesNumber; i++ {
		fmt.Println(accounts[i])
	}
	fmt.Println(cp.ConcernPlatesID)
	// fmt.Println(cp.ConcernPlatesNumber)
	//1471290040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
	//147129004000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002
}

package main

import (
	"fmt"
	"net/rpc"
)

type Args struct { //只要c端的有属性(名字和类型)和s端相同即可
	Num1 int32
	Num2 int32
	sum  int32
}

func main() {
	cal, err := rpc.DialHTTP("tcp", "127.0.0.1:8080") //链接s端
	if err != nil {
		fmt.Println("error")
	}

	args := new(Args)
	args.Num1 = 8
	args.Num2 = 6

	err = cal.Call("Calculate.Add", args, &args.sum) //调用
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("sum=", args.sum)
}

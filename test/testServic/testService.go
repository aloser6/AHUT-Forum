package main

import (
	services "ISPS/web/service/src"
	"ISPS/web/utils"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	//testTime()
	//testRpcServer()
	//testRpcClient()
	testService()
}

func testTime() {
	t := new(utils.Timer)
	var num *int8
	t.AddTimerTask(&utils.TimerTask{Time: 3, RealTime: 0, Task: func() { fmt.Fprintln(os.Stdout, "3 second") }}, num)
	t.AddTimerTask(&utils.TimerTask{Time: 2, RealTime: 0, Task: func() { fmt.Fprintln(os.Stdout, "2 second") }}, num)
	t.ExecTimer("123")
}

type Calculcate struct {
}
type Args struct {
	Num1 int32
	Num2 int32
}

func (cal *Calculcate) Add(args *Args, rep *int32) error {
	*rep = args.Num1 + args.Num2
	return nil
}

func testRpcServer() {
	rpcs := new(utils.Rpcer)
	rpcs.Register(new(Calculcate))
	l, e := net.Listen("tcp", "127.0.0.1:8080")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func testRpcClient() {
	var val int32
	rpcs := new(utils.Rpcer)

	err := rpcs.Call("127.0.0.1:8080", "Calculcate.Add", Args{6, 4}, &val)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

func testService() {
	var val *int8
	rpcs := new(utils.Rpcer)
	err := rpcs.Call("127.0.0.1:8080", "Centre.Register", services.Service{Sname: "123", Ip: "127.0.0.1", Port: 8080}, val)
	if err != nil {
		fmt.Println(err)
	}
	times := new(utils.Timer)
	times.ExecTimer("123")
}

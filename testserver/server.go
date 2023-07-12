package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	cal := new(Calculate)
	rpc.Register(cal)                           //注册
	rpc.HandleHTTP()                            // 将 RPC 服务绑定到 HTTP 服务中去
	l, e := net.Listen("tcp", "127.0.0.1:8080") //创建监听
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil) //服务
	//大小端问题
}

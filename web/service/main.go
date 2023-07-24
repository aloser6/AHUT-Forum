package main

import (
	"ISPS/web/utils"
	"log"
	"net"
	"net/http"

	services "ISPS/web/service/src"
)

func main() {
	initService()
}

func initService() {
	rpcs := new(utils.Rpcer)
	rpcs.Register(new(services.Centre))
	for {
		l, e := net.Listen("tcp", "127.0.0.1:8080")
		if e != nil {
			log.Fatal("listen error:", e)
		}
		http.Serve(l, nil)
	}
}

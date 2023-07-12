package main

type Args struct {
	Num1 int32
	Num2 int32
}

type Calculate struct {
}

func (cal *Calculate) Add(args *Args, rep *int32) error {
	*rep = args.Num1 + args.Num2
	return nil
}
func (cal *Calculate) Sub(args *Args, rep *int32) error {
	*rep = args.Num1 - args.Num2
	return nil
}
func (cal *Calculate) Mul(args *Args, rep *int32) error {
	*rep = args.Num1 * args.Num2
	return nil
}
func (cal *Calculate) Div(args *Args, rep *int32) error {
	*rep = args.Num1 / args.Num2
	return nil
}

// package main

// import (
// 	"net"
// 	"net/http"
// 	"net/rpc"
// )

// type Calculate struct {
// 	method string
// }
// type Args struct {
// 	Num1 int32
// 	Num2 int32
// 	Num3 int32
// 	Ans  int32
// }

// func main() {
// 	cal := new(Calculate)
// 	rpc.Register(cal) //注册
// 	rpc.HandleHTTP()  // ??
// 	lis, err := net.Listen("tcp", "127.0.0.1:8080")
// 	if err != nil {
// 		return
// 	}
// 	go http.Serve(lis, nil)
// }

// func (cal *Calculate) Methods(method string, args *Args) error {
// 	cal.method = method
// 	switch cal.method {
// 	case "+":
// 		cal.Add(args)
// 	case "-":
// 		cal.Sub(args)
// 	case "*":
// 		cal.Mul(args)
// 	case "/":
// 		cal.Div(args)
// 	default:
// 		return nil //TODO
// 	}
// 	return nil
// }

// func (cal *Calculate) Add(args *Args) {
// 	args.Ans = args.Num1 + args.Num2 + args.Num3
// }
// func (cal *Calculate) Sub(args *Args) {
// 	args.Ans = args.Num1 - args.Num2 - args.Num3
// }
// func (cal *Calculate) Mul(args *Args) {
// 	args.Ans = args.Num1 * args.Num2 * args.Num3
// }
// func (cal *Calculate) Div(args *Args) {
// 	args.Ans = args.Num1 / args.Num2 / args.Num3
// }

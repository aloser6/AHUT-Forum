package utils

import (
	"errors"
	"net/rpc"
)

type Rpcer struct {
	//client *rpc.Client
}

/*功能:实现对结构体方法的注册(仅注册，未监听)
 *参1:要注册的服务的结构体对象
 **/
func (r *Rpcer) Register(class interface{}) error {
	if class == nil {
		return errors.New("invalid args")
	}
	err := rpc.Register(class)
	if err != nil {
		return err
	}
	rpc.HandleHTTP()

	return nil
}

//监听
/* l, e := net.Listen("tcp", ":1234")
if e != nil {
	log.Fatal("listen error:", e)
}
go http.Serve(l, nil)
*/

/*功能:实现函数的远程调用
 *参1: ip:端口号 的字符串
 *参2:函数名
 *参3:参数
 *参4:返回值(使用指针/结构体指针)
 **/
func (r *Rpcer) Call(iport string, funName string, args interface{}, val interface{}) error {
	if iport == "" || funName == "" {
		return errors.New("invalid args")
	}

	client, err := rpc.DialHTTP("tcp", iport)
	if err != nil {
		return err
	}
	err = client.Call(funName, args, val)
	if err != nil {
		return err
	}
	//defer client.Close()??
	return nil
}

//TODO参数主机序转网络序

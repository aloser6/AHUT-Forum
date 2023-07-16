package utils

import (
	"errors"
	"net/rpc"
)

type Rpcer struct {
	//client *rpc.Client
}

/*功能:实现对结构体方法的注册(仅注册，未监听)
 *参1:结构体对象
 *参2:ip:port的字符串
 **/
func (r *Rpcer) Register(class interface{}) error {
	if class == nil {
		return errors.New("invalid args")
	}

	if err := rpc.Register(class); err != nil {
		return err
	}
	rpc.HandleHTTP()

	return nil
}

/*功能:实现函数的远程调用
 *参1: ip:端口号 的字符串
 *参2:函数名
 *参3:返回值(使用指针/结构体指针)
 *参4:参数(变长)
 **/
func (r *Rpcer) Call(iport string, funName string, args interface{}, val interface{}) error {
	if iport == "" || funName == "" {
		return errors.New("invalid args")
	}

	client, err := rpc.Dial("tcp", iport)
	if err != nil {
		return errors.New("rpc.Dial fail")
	}
	err = client.Call(funName, args, val)
	if err != nil {
		return errors.New("rpc.Call fail")
	}
	//defer client.Close()??
	return nil
}

//TODO参数主机序转网络序

package service

import (
	"errors"
	"sync"
)

type Service struct {
	Sname string //服务器名
	Ip    string //ip
	Port  int32  //port
	//TODO time
	LiveTime int64 `default:""`      //心跳时间
	IsLive   bool  `default:"false"` //是否存活
}

type Centre struct {
	Service []Service      //存储服务对象
	Index   map[string]int //Service的索引s
	Mutex   sync.Mutex     //锁
	IsLock  bool           `default:"false"` //是否已锁
}

/*ipv4
 *功能:服务注册
 *参1:存储服务信息的结构体
 *参2:任意*int16的值
 **/
func (c *Centre) RegisterLock(args Service, val *int16) error { //参数优化val
	if args.Ip == "" || args.Port < 0 {
		return errors.New("invalid args")
	}
	if args.IsLive == true {
		return errors.New("already registered service")
	}

	//TODO读写锁
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.IsLock = true

	c.Service = append(c.Service, args)
	if _, ok := c.Index[args.Sname]; ok == true {
		c.IsLock = false
		return errors.New("name repeat")
	}

	c.Index[args.Sname] = len(c.Service) - 1
	c.Service[c.Index[args.Sname]].IsLive = true
	c.IsLock = false

	return nil
}

func (c *Centre) Discover() { //读写锁

}
func (c *Centre) AddTimeLock() {

}
func (c *Centre) DeleteLock() {

}

//ipv6

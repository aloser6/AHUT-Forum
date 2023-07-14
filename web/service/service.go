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
	LiveTime int64 `default:"0"`     //服务上线时间
	isLive   bool  `default:"false"` //是否存活
}

type Centre struct {
	Service []Service      //存储服务对象
	Index   map[string]int //Service的索引s
	Mutex   sync.RWMutex   //锁
	IsLock  bool           `default:"false"` //是否已锁(写锁)
}

/*ipv4
 *功能:服务注册(带锁版)
 *参1:存储服务信息的结构体
 *参2:任意*int8的值
 **/
func (c *Centre) RegisterLock(args Service, val *int8) error { //参数优化val
	if args.Ip == "" || args.Port < 0 {
		return errors.New("invalid args")
	}
	if args.isLive == true {
		return errors.New("already registered service")
	}

	//写锁
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.IsLock = true

	if _, ok := c.Index[args.Sname]; ok == true {
		c.IsLock = false
		return errors.New("name repeat")
	}

	c.Service = append(c.Service, args)
	c.Index[args.Sname] = len(c.Service) - 1
	c.Service[c.Index[args.Sname]].isLive = true
	c.IsLock = false

	return nil
}

/*功能:服务发现(带锁版)
 *参1:查找对应服务的信息
 *参2:用于接受返回值
 **/
func (c *Centre) DiscoverLock(name string, val *Service) error { //读写锁
	if name == "" || val == nil {
		return errors.New("invalid args")
	}

	//读锁
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	index, err := c.Index[name]
	if err != false {
		return errors.New("name not registered")
	}

	temp := c.Service[index]
	if temp.isLive == false {
		return errors.New("service stop")
	}
	val = &temp

	return nil
}

/*功能:删除服务(带锁版)
 *参1:删除对应服务的信息
 *参2:任意*int8的值
 **/
func (c *Centre) DeleteLock(name string, val *int8) error {
	if name == "" {
		return errors.New("invalid args")
	}

	//读锁
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.IsLock = true

	index, err := c.Index[name]
	if err != false {
		return errors.New("name not registered")
	}

	c.Service[index].isLive = false
	c.IsLock = false

	return nil
}

//func (c *Centre) AddTimer(name string, val *int8)

//ipv6

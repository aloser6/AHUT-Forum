package utils

import (
	"container/list"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

type TimerTask struct {
	Time     int64  `default:"1e9"` //倒计时的时间
	RealTime int64  //相对时间
	Task     func() //到时间后需要执行的任务(回调函数)
}

type Timer struct {
	TimeList     *list.List   //存储定时器
	Mutex        sync.RWMutex //读写锁
	hertBeatTime int32        `default:"1e9"`
}

/*功能:添加定时任务
 *参1:定时任务结构体
 *参2:任意*int8的值
 **/
func (t *Timer) AddTimerTask(args *TimerTask, val *int8) error {
	if args == nil {
		return errors.New("invalid args")
	}

	// task := new(TimerTask)
	// task.Time = args.Time
	// task.Task = args.Task
	// task.RealTime = 0

	err := t.insertTask(args)
	if err != nil {
		return err
	}

	return nil
}

// 暂且只支持最小粒度为1s的心跳，待突破
func (t *Timer) insertTask(task *TimerTask) error { //可算法优化
	if task == nil {
		return errors.New("invalid args")
	}
	if task != nil && task.Time < 0 {
		go task.Task()
	}

	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	//开头
	if t.TimeList == nil {
		t.TimeList = new(list.List)
		t.TimeList.Init()
		task.RealTime = task.Time
		t.TimeList.PushBack(task)
		return nil
	} else if t.TimeList.Front() == nil {
		task.RealTime = task.Time
		t.TimeList.PushBack(task)
		return nil
	}

	//中间
	var last int64 = 0
	for i := t.TimeList.Front(); i != t.TimeList.Back().Next(); i = i.Next() {
		current, ok := i.Value.(*TimerTask)
		if !ok {
			return errors.New("type error")
		}
		task.RealTime = task.Time - last
		if task.RealTime < current.RealTime {
			current.RealTime -= task.RealTime
			t.TimeList.InsertBefore(task, i)
			return nil
		}
		last += current.RealTime
	}

	//结尾
	task.RealTime = task.Time - last
	t.TimeList.PushBack(task)
	return nil
}

/*功能:执行定时器
 *参1:服务名称
 **/
func (t *Timer) ExecTimer(serviceName string) error {
	timer := time.NewTimer(time.Duration(t.hertBeatTime))
	//rpcs := Rpcer{}
	var val int8
	if t.TimeList == nil {
		t.TimeList = new(list.List)
	}
	t.AddTimerTask(&TimerTask{1, 0, func() {
		fmt.Fprintln(os.Stdout, "1 second1") /*rpcs.Call("127.0.0.1:8080", "Centre.HertBeat", serviceName, &val)*/
	}}, &val) //ip:port等配置系统优化

	for {
		nextTime, ok := t.TimeList.Front().Value.(*TimerTask)
		if !ok {
			return errors.New("type error")
		}
		timer.Reset(time.Duration(nextTime.RealTime * int64(time.Second)))
		<-timer.C //等1s
		fmt.Println(".")

		t.Mutex.Lock() //如果等待过长时间怎么办？？
		temp, ok := t.TimeList.Front().Value.(*TimerTask)
		if !ok {
			return errors.New("type error")
		}
		temp.RealTime -= 1

		i := t.TimeList.Front()
		for i != nil {
			task, ok := i.Value.(*TimerTask)
			if !ok {
				return errors.New("type error")
			}
			if task.RealTime <= 0 { //执行的是剩1s的
				temp := i
				i = i.Next()
				t.TimeList.Remove(temp)
				if task.Task != nil {
					go task.Task()
				}
				continue
			}
			break
		}

		t.Mutex.Unlock()
		t.AddTimerTask(&TimerTask{1, 0, func() {
			fmt.Fprintln(os.Stdout, "1 second2") /*rpcs.Call("127.0.0.1:8080", "Centre.HertBeat", serviceName, &val)*/
		}}, &val) //ip:port等配置系统优化

	}
}

// 待突破部分
// if i == t.TimeList.Front() {
// 	times, err := t.SignalTime() //insertFront()//current.RealTime -= //已经运行的时间
// 	if err != nil {s
// 		return nil, errors.New("SignalTime fail")
// 	}
// 	current.RealTime -= times
// }

//fmt.Println(t.TimeList.Len())
// for i := t.TimeList.Front(); i != t.TimeList.Back().Next(); { //执行所有超时任务
// 	task, ok := i.Value.(*TimerTask)
// 	if !ok {
// 		return errors.New("type error")
// 	}
// 	fmt.Println(task.RealTime)
// 	if task.RealTime <= 0 { //执行的是剩1s的
// 		temp := i
// 		i = i.Next()
// 		t.TimeList.Remove(temp)
// 		if task.Task != nil {
// 			go task.Task()
// 		}
// 		for s := t.TimeList.Front(); s != t.TimeList.Back().Next(); s = s.Next() {
// 			tt := s.Value.(*TimerTask)
// 			fmt.Println("s:", tt.Time)
// 		}
// 		continue
// 	}

// 	break
// }
//fmt.Println(t.TimeList.Len())

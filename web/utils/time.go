package utils

import (
	"container/list"
	"errors"
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

	task := new(TimerTask)
	task.Time = args.Time
	task.Task = args.Task
	task.RealTime = 0

	index, err := t.findLowerBound(task)
	if err != nil {
		return errors.New("FindLowerBoundLock fail")
	}
	if index != nil {
		t.TimeList.InsertBefore(task, index)
	}

	return nil
}

// 暂且只支持最小粒度为1s的心跳，待突破
func (t *Timer) findLowerBound(task *TimerTask) (*list.Element, error) { //可算法优化
	if task != nil && task.Time < 0 {
		go task.Task()
	}

	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	//开头
	if t.TimeList.Front() == nil {
		t.TimeList.Init()
		task.RealTime = task.Time
		t.TimeList.PushBack(task)
		return nil, nil
	}

	//中间
	var last int64 = 0
	for i := t.TimeList.Front(); i != t.TimeList.Back(); i = i.Next() {
		current, ok := i.Value.(*TimerTask)
		if ok {
			return nil, errors.New("type error")
		}
		task.RealTime = task.Time - last
		if task.RealTime < current.RealTime {
			current.RealTime -= task.RealTime
			return i, nil
		}
		last += current.RealTime
	}

	//结尾
	task.RealTime = task.Time - last
	t.TimeList.PushBack(task)
	return nil, nil
}

/*功能:执行定时器
 *参1:服务名称
 **/
func (t *Timer) ExecTimer(serviceName string) error {
	timer := time.NewTimer(time.Duration(t.hertBeatTime))
	for {
		nextTime := t.TimeList.Front().Value.(TimerTask)
		timer.Reset(time.Duration(nextTime.RealTime))
		<-timer.C //等1s

		t.Mutex.Lock() //如果等待过长时间怎么办？？
		task := t.TimeList.Front().Value.(TimerTask)
		t.TimeList.Remove(t.TimeList.Front())
		rpcs := Rpcer{}
		var val int8
		t.AddTimerTask(&TimerTask{1, 0, func() { rpcs.Call("127.0.0.1:8080", "Centre.HertBeat", serviceName, &val) }}, &val) //ip:port等配置系统优化
		t.Mutex.Unlock()

		go task.Task()
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

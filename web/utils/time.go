package utils

import (
	"container/list"
	"errors"
	"sync"
)

type TimerTask struct {
	Time     int32  //倒计时的时间
	RealTime int32  //相对时间
	Task     func() //到时间后需要执行的任务(回调函数)
}

type Timer struct {
	TimeList *list.List   //存储定时器
	Mutex    sync.RWMutex //读写锁
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
		//TODO执行
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
	var last int32 = 0
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

// func (t *Timer) ExecTimer() (int32, error) {
// }

// 待突破部分
// if i == t.TimeList.Front() {
// 	times, err := t.SignalTime() //insertFront()//current.RealTime -= //已经运行的时间
// 	if err != nil {
// 		return nil, errors.New("SignalTime fail")
// 	}
// 	current.RealTime -= times
// }

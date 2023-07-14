package utils

import (
	"container/list"
	"errors"
)

type TimerTask struct {
	Time     int32  //倒计时的时间
	RealTime int32  //相对时间
	Task     func() //到时间后需要执行的任务(回调函数)
}

type Timer struct {
	TimeList *list.List //存储定时器
}

func (t *Timer) AddTimerTaskLock(args *TimerTask, val *int8) error {
	if args == nil {
		return errors.New("invalid args")
	}

	task := new(TimerTask)
	task.Time = args.Time
	task.Task = args.Task
	task.RealTime = 0

	//TODO insert

	return nil
}

// check
func (t *Timer) FindLowerBoundLock(task *TimerTask) (*TimerTask, error) { //可算法优化
	//TODO lock
	if task != nil && task.Time < 0 {
		//TODO执行
	}
	if t.TimeList.Front() == nil {
		t.TimeList.Init()
		task.RealTime = task.Time
		t.TimeList.PushBack(task)
		return task, nil //TODO在开头 //task前插入task
	}

	var last int32 = 0
	for i := t.TimeList.Front(); i != t.TimeList.Back(); i = i.Next() {
		current, ok := i.Value.(*TimerTask)
		task.RealTime = task.Time - last

		if ok {
			return nil, errors.New("type error")
		}
		if task.RealTime <= current.RealTime {
			return current, nil //TODO在中间
		}
	}
	return nil, nil //TODO末尾
}

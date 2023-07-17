package main

import (
	"ISPS/web/utils"
	"fmt"
	"os"
)

func main() {
	t := new(utils.Timer)
	var num *int8
	t.AddTimerTask(&utils.TimerTask{Time: 3, RealTime: 0, Task: func() { fmt.Fprintln(os.Stdout, "3 second") }}, num)
	t.AddTimerTask(&utils.TimerTask{Time: 2, RealTime: 0, Task: func() { fmt.Fprintln(os.Stdout, "2 second") }}, num)
	t.ExecTimer("123")

}

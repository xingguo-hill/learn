package x

import (
	"fmt"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	deleySecond := 2
	t1 := time.Now()
	time.Sleep(time.Second * time.Duration(deleySecond))
	duration := time.Now().Sub(t1)
	fmt.Printf("sleep %d秒，历时%d秒\n", deleySecond, int(duration.Seconds()))
}
func TestTimer(t *testing.T) {
	i := 0
	d := time.Duration(1) * time.Second
	timer := time.NewTimer(d)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			i++
			//每秒执行1次
			fmt.Printf("timer 1s %d\n", i)
			//循环与internal类似
			timer.Reset(d)
		}
	}
}

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

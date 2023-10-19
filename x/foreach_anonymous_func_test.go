package x

import (
	"fmt"
	"testing"
	"time"
)

func TestForeach(t *testing.T) {
	subs := []int{1, 2, 3}
	for _, v := range subs {
		go func(i int) {
			fmt.Println(i)
		}(v)
	}
	time.Sleep(time.Second)
}

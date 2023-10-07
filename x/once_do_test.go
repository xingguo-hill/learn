package x

import (
	"fmt"
	"sync"
	"testing"
)

var once sync.Once

func TestOnceDo(t *testing.T) {
	for i := 0; i < 3; i++ {
		once.Do(func() {
			fmt.Println("我只执行1次")
		})
	}

}

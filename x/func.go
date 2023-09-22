package x

import (
	"fmt"
	"math"
	"time"
)

func Run() {
	// stopChildByChannel()
	// useIota()
	DeferClosure()
}

func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}

func DeferClosureV1() {
	i := 0
	defer func(val int) {
		println(val)
	}(i)
	i = 1
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func YourName(name string, alias ...string) {
	if len(alias) > 0 {
		fmt.Printf("%#v", alias)
		println(len(alias))
	}
}

func YourNameInvoke() {
	YourName := func(name string, alias ...string) {
		if len(alias) > 0 {
			fmt.Printf("%#v", alias)
			println(len(alias))
		}
	}
	YourName("Deng Ming")
	YourName("Deng Ming", "Da Ming")
	YourName("Deng Ming", "Da Ming", "Zhong Ming")
}
func StopChildByChannel() {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done channel is triggerred, exit child go routine")
				return
			}
		}
	}()
	close(done)
	time.Sleep(time.Second)
}
func UseIota() {
	// var a int8.maximum = -128
	fmt.Printf("\"hello,world")
	fmt.Println(math.Abs(123.2))
	a := 1
	fmt.Println(a)
	go func() {
		fmt.Println(a)
	}()
	const (
		start = iota << 1
		run
		down
		err
	)
	fmt.Println(start, run, down, err)
}

package x

import (
	"fmt"
	"testing"
	"time"
)

// 定时器的使用
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

// 时间的差值计算
func TestDuration(t *testing.T) {
	deleySecond := 2
	t1 := time.Now()
	time.Sleep(time.Second * time.Duration(deleySecond))
	duration := time.Now().Sub(t1)
	fmt.Printf("sleep %d秒，历时%d秒\n", deleySecond, int(duration.Seconds()))
}

// 时间对象格式化输出
func TestTimeFormat(t *testing.T) {
	t1 := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(t1)
	t1 = time.Now().Format("2006-01-02")
	fmt.Println(t1)
}

// 指定日期转化为时间对象
func TestDateToTime(t *testing.T) {
	// 指定的日期字符串
	dateString := "2023-10-30"
	// 定义日期和时间的格式
	layout := "2006-01-02"
	// 将日期字符串解析为时间对象
	parsedTime, _ := time.Parse(layout, dateString)
	// 输出解析后的时间对象
	fmt.Println(parsedTime)
}

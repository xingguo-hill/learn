package x

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	// 创建一个根上下文
	ctx := context.Background()

	// 创建一个具有取消特性的上下文，使用 context.WithCancel
	ctx, cancel := context.WithCancel(ctx)

	// 启动一个并发操作
	go performTask(ctx)
	time.Sleep(time.Second)
	// 执行取消操作
	cancel()
	// 模拟等待一段时间
	time.Sleep(3 * time.Second)
}

func performTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// 检查上下文的 Done 通道
			fmt.Println("Task canceled or timed out.")
			return
		default:
			// 执行任务
			fmt.Println("Performing task...")
			time.Sleep(2 * time.Second)
		}
	}
}

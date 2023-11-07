package x

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSignal(t *testing.T) {
	// context init
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	chSig := make(chan os.Signal, 1)
	/* kill -xxx pid
	* syscall.SIGINT  值 2 用户发送INTR字符(Ctrl+C)触发
	* syscall.SIGQUIT 值 3 用户发送QUIT字符(Ctrl+\)触发
	* syscall.SIGTERM 值 15 结束程序(可以被捕获、阻塞或忽略)
	 */
	signal.Notify(chSig, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	logrus.Println("process ID:", os.Getpid())
	go func() {
		logrus.Println("parent process ID:", os.Getppid())
		logrus.Println("current process ID:", os.Getpid())
		<-ctx.Done()
		logrus.Println("parent process return:")
		return
	}()
	sig := <-chSig
	cancel() // notify goroutines close
	logrus.Println("process quit... Signal value", sig)
}

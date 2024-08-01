package x

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"
)

func TestPipe(t *testing.T) {
	// 创建管道
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println("Error creating pipe:", err)
		return
	}

	// 创建子进程并将标准输出重定向到管道的写入端
	cmd := exec.Command("echo", "Hello from the child process!")
	cmd.Stdout = w

	// 启动子进程
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	// 关闭写入端，因为子进程已经结束
	w.Close()

	// 从读取端读取子进程的输出
	buffer := make([]byte, 1024)
	n, err := r.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading from pipe:", err)
		return
	}

	// 输出子进程的输出
	fmt.Printf("Child process output: %s\n", buffer[:n])
	// 关闭读取端
	r.Close()
}

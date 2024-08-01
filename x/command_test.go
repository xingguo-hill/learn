package x

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
)

func cmdNoOutput(ip string) error {
	cmd := exec.Command("ping", ip, "-c", "1", "-t", "1")
	return cmd.Run()
}
func cmdOutput() bool {
	cmd := exec.Command("go", "version")
	out, _ := cmd.Output()
	return strings.Contains(cast.ToString(out), "version")
}
func TestCmd(t *testing.T) {
	assert.Equal(t, nil, cmdNoOutput("180.76.76.76"))
	assert.Equal(t, true, cmdOutput())
}
func TestCmdInput(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println("Error creating pipe:", err)
		return
	}
	go func() {
		defer w.Close()
		fmt.Fprintln(w, "Hello from the findnotme!")
		fmt.Fprintln(w, "Hello from the findme!")
	}()

	cmd := exec.Command("grep", "findme") // grep 命令会从标准输入中读取并查找包含 "findme" 的行

	// 将标准输入传递给子进程
	cmd.Stdin = r

	// 执行子进程
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印子进程的输出
	fmt.Println("Output from child process:")
	fmt.Println(string(output))
}

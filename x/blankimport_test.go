package x

import (
	"fmt"
	_ "learn/x/blankimport"
	"testing"
)

// 测试空白符导入，只支持init函数执行
func Test_init(t *testing.T) {
	fmt.Println("test end")
}

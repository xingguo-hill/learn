package x

import (
	"fmt"
	"testing"
)

type I struct {
	NoUse int
}

type Cmdable interface {
	A(int) int
}

func (i *I) A(a int) int {
	fmt.Println(a)
	return a
}
func (i *I) B(b string) string {
	fmt.Println(b)
	return b
}

func YourTest() Cmdable {
	i := &I{}
	i.A(1)
	return i
}

func TestYourTest(t *testing.T) {
	YourTest()
	t.Log("测试中!")
}

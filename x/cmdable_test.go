package x

import (
	"fmt"
	"testing"
)

type I struct {
	NoUse int
}

type Cmdable interface {
	A(int)
}

func (i *I) A(a int) {
	fmt.Println(a)
}
func (i *I) B(b string) string {
	fmt.Println(b)
	return b
}

func YourTest() Cmdable {
	i := &I{NoUse: 1}
	i.A(1)
	return i
}

func TestYourTest(t *testing.T) {
	//因此发生隐式转换，ti为接口，不再是结构体I
	ti := YourTest()
	fmt.Printf("%#v", ti)
	t.Log("测试中!")
}

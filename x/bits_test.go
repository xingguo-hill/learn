package x

import (
	"fmt"
	"testing"
)

func TestBit(t *testing.T) {
	//占4位
	code := 15
	bits := fmt.Sprintf("%016b", uint16(code)<<11)
	fmt.Println(bits)

	//占最高1位
	code = 0
	_QR := 1 << 15
	code |= _QR
	bits = fmt.Sprintf("%016b", code)
	fmt.Println(bits)

}

package x

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestUtf8(t *testing.T) {
	// 使用 utf8mb4 编码的字符串，包括 emoji 字符
	str := "你好H 🚀"

	// 使用 utf8.RuneCountInString 计算字符串的长度
	length := utf8.RuneCountInString(str)

	fmt.Printf("String: %s\n", str)
	fmt.Printf("Length (in runes): %d\n", length)
}

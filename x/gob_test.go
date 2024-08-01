package x

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

func TestGob(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	/** ------------------编码------------------*/
	// 创建要编码的对象
	person := Person{Name: "Alice", Age: 30}

	// 创建一个缓冲区用于保存编码后的数据
	var ebuf bytes.Buffer

	// 创建编码器，编码对象
	err := gob.NewEncoder(&ebuf).Encode(&person)
	if err != nil {
		fmt.Println("编码错误:", err)
	}
	// 编码后的数据存储在 buf 中
	fmt.Println("编码后的数据:", ebuf.Bytes())

	/** ------------------解码------------------*/
	// 创建缓冲区并写入数据
	dbuf := bytes.NewBuffer(ebuf.Bytes())
	// 创建解码器
	dec := gob.NewDecoder(dbuf)

	// 解码数据
	err = dec.Decode(&person)
	if err != nil {
		fmt.Println("解码错误:", err)
	}

	// 输出解码后的对象
	fmt.Printf("解码后的对象:%#v", person)
}

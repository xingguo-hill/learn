package blankimport

import "fmt"

type BI struct {
	T string
}

var B BI

func init() {
	B.T = ""
	fmt.Println("init() run")
}

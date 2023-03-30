package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 要实现的效果 （尾行达 25 个字符为止）
	G
	GG
	GGG
	GGGG
	GGGGG
	GGGGGG
	*/
	// 2层for循环
	for i := 1; i <= 25; i++ {
		fmt.Println(strings.Repeat("G", i))
	}
}
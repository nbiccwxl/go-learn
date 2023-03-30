package main

import (
	"fmt"
)

func main() {
	initValue := 10
	toPrint(initValue)
}

func toPrint(value int) {
	fmt.Printf("输出值：%d\n", value)
	value--
	if value >= 1 {
		toPrint(value)
	}
}





// 使用递归函数从 10 打印到 1。


package main

import (
	"fmt"
)

func main() {
	num1, num2, num3 := numText(4,3)
	fmt.Printf("非命名返回值：%d, %d, %d", num1, num2, num3)
	num1, num2, num3 = numText2(4,3)
	fmt.Printf("非命名返回值：%d, %d, %d", num1, num2, num3)
}

// 编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。

func numText(x1, x2 int) (int, int, int) {
	return x1 + x2, x1 * x2, x1 - x2
}

func numText2(x1, x2 int) (f1 int, f2 int, f3 int) {
	f1 = x1 + x2
	f2 = x1 * x2
	f3 = x1 - x2
	return
}
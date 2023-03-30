package main

import (
	"fmt"
)

func main() {
	// 声明切片的格式
	// var identifier []type

	// 切片的初始化格式 ： start:end 切片表达式 slice1[0] = arr1[start]
	// var slice1 []type = arr1[start:end]
	fmt.Println("gg")
	arr := [3]int{1, 2, 3}
	s := arr[:len(arr)]
	x := arr[0:]
	var b bool = equal(s, x)
	fmt.Printf("s,x的比较结果为：%v, %t, %v", s, b, x)
	fmt.Printf("slice: %d", s)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

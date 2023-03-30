package main

import "fmt"

func main() {
    i := 100

    addNum := func(j int) int { //定义一个闭包函数
        i = i + j
        return i
    }

	

    fmt.Println(addNum(1)) // 输出 101
	i = 102;
    fmt.Println(addNum(2)) // 输出 103
    fmt.Println(addNum(3)) // 输出 106
}
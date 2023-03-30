package main

import (
	"fmt"
	"math"
)

func main() {
	result, err := MySqrt2(9)
	if err != nil {
		fmt.Printf("发生错误", err)
	} else {
		fmt.Printf("计算完成：", result)
	}
}

// 编写一个名字为 MySqrt() 的函数，计算一个 float64 类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。编写两个版本，一个是非命名返回值，一个是命名返回值。

func MySqrt(num float64) (result float64, err error) {
	if num < 0 {
		err = fmt.Errorf("出现负数")
		return
	}
	result = math.Sqrt(num)
	fmt.Println("结果：%d", result)
	return
}

func MySqrt2(num float64) (float64, error) {
	if num < 0 {
		err := fmt.Errorf("出现负数")
		return 0.0, err
	}
	result := math.Sqrt(num)
	fmt.Println("结果：%d", result)
	return result, nil
}
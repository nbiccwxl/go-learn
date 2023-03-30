package main

import "fmt"

func main() {
	for i := 0; i <= 15; i++ {
		fmt.Printf("for循环下数字位：%d\n", i)
	}
	sex := 0;
	loop:
	if sex < 5 {
		sex += 1;
		fmt.Printf("不满足sex条件，当前sex：%d\n", sex)
		goto loop;
	} else {
		fmt.Printf("goto循环下数字位：%d\n", sex)
	}

}
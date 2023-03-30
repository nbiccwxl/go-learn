package main
import "fmt"

type TZ int

type Proe string

func main() {
	var a, b TZ = 3, 4
	c := a + b
	var d, e Proe = "a", "b"
	fmt.Printf("数据如下所示, d: %v, e: %v", d, e)
	fmt.Println()
	fmt.Printf("c has the value: %d hh", c) // 输出：c has the value: 7
}
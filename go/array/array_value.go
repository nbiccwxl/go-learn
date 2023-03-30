package main
import "fmt"

func main() {
	var arr [5]int
	fmt.Println("first main方法中的：", &arr)
	arrayPoint(arr)
	fmt.Println("second main方法中的：", &arr)
}


func arrayPoint(arr [5]int) {
	arr[2] = 4
	fmt.Println("func中的地址：", &arr)
}
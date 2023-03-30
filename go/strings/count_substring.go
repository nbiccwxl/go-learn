package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str)
	// 2
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	// 5
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
	fmt.Println()
	fmt.Printf("new string is %s", strings.Replace(manyG, "g", "s", -1))
}
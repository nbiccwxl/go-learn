package main
import "fmt"
func main() {
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	// s 的内存地址
	fmt.Printf("Here is the pointer p: %p\n", p) // prints address
	// ciao
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	// ciao
	fmt.Printf("Here is the string s: %s\n", s) // prints same string
}
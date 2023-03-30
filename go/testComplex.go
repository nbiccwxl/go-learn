package main


import "fmt"

func main() {
    re := float32(3.14)
    im := float32(2.78)
    c := complex(re, im)
    fmt.Println(c)
}
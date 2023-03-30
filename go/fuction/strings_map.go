package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "Hello, 您好！This is a test 1234."
	f := func(c rune) bool {
		return !unicode.Is(unicode.ASCII_Hex_Digit, c)
	}
	replacer := strings.NewReplacerFunc(f, func(c rune) string {
		if unicode.IsSpace(c) {
			return " "
		} else {
			return "?"
		}
	})
	result := replacer.Replace(text)
	fmt.Println(result)
}

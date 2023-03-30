package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i <= 100; i++ {
		
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
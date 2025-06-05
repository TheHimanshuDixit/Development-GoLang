package main

import (
	"fmt"
)

func main() {
	num := 10
	// switch statement
	switch {
	case num%2 == 0:
		fmt.Println(num, "is even")
	default:
		fmt.Println(num, "is odd")
	}
}

package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var y int = 20
	var ptr *int

	ptr = &x
	fmt.Println("Value of x:", *ptr)

	ptr = &y
	fmt.Println("Value of y:", *ptr)

	num := 30
	ptr = &num
	fmt.Println("Value of num:", *ptr)

	// rest all are like cpp
}

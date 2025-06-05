package main

import (
	"fmt"
)

func main() {
	PrintMessage()
	PrintNumber(42)
}

func PrintMessage() {
	fmt.Println("Hello, World!")
}

func PrintNumber(num int) {
	fmt.Println("The number is:", num)
	fmt.Printf("The number is: %d\n", num)

	// %T is for type
	fmt.Printf("The type of num is: %T\n", num)
}

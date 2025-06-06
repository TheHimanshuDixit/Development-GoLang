package main
import (
	"fmt"
)

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	// Output:
	// Hello
	// World

	// double defer
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	// Output:
	// Second defer
	// First defer
}
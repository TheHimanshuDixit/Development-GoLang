package main

import (
	"fmt"
)

func main() {
	// Define an array of integers
	numbers := [5]int{1, 2, 3, 4, 5}

	// Print the array
	fmt.Println("Array of integers:", numbers)
	// Length of the array
	fmt.Printf("Length of the array: %d\n", len(numbers))

	names := [3]string{"Alice", "Bob", "Charlie"}
	// Print the array of strings
	fmt.Println("Array of strings:", names)

	var num [3]int
	fmt.Println("Array of integers with uninitialized values:", num)

	var price [3]string
	fmt.Println("Array of strings with uninitialized values:", price)
	fmt.Printf("Length of the price array: %q\n", price)
}

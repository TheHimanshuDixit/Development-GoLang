package main

import (
	"fmt"
)

func main() {
	// Define a slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Call the function to calculate the sum
	sum := calculateSum(numbers)

	// Print the result
	fmt.Println("The sum of the numbers is:", sum)

	// Example of a function with named return values
	result := add(10, 20)
	fmt.Println("The result of adding 10 and 20 is:", result)
}

func add(a, b int) (result int) {
	result = a + b
	return
}

func calculateSum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

package main

import (
	"fmt"
)

func main() {

	num := []int{1, 2, 3, 4, 5}
	num = append(num, 6)       // Appending a new element to the slice
	num = append(num, 7, 8, 9) // Appending multiple elements to the slice
	fmt.Println("Slice of integers:", num)
	fmt.Printf("Length of the slice: %d\n", len(num))

	numbers := make([]int, 5, 10) // Creating a slice with a length of 5 and a capacity of 10
	fmt.Println("Slice of integers with length 5 and capacity 10:", numbers)
}

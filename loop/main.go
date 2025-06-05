package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// infinite loop example
	// for {
	// 	fmt.Println("This will run forever unless stopped.")
	// Uncomment the next line to break the infinite loop after one iteration
	// break
	// }

	// loop using range
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	data := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range data {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	strinng := "hello"
	for index, char := range strinng {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}

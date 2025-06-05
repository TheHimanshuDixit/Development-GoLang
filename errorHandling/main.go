package main

import (
	"fmt"
)

func main() {
	result, err := divide(10, 0)
	// result, _ := divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The result of the division is:", result)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}


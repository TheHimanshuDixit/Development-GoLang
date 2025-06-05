package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 10
	var b *int = &a                        // b is a pointer to a
	fmt.Println("Value of a:", a)          // Output: Value of a: 10
	fmt.Println("Address of a:", &a)       // Output: Address of a: <memory address>
	fmt.Println("Value of b:", b)          // Output: Value of b: <memory address>
	fmt.Println("Value pointed by b:", *b) // Output: Value pointed by b: 10

	*b = 20                           // Change the value at the address pointed by b
	fmt.Println("New value of a:", a) // Output: New value of a: 20

	var num int = 30
	var data float64 = float64(num)       // Convert int to float64
	fmt.Println("Converted value:", data) // Output: Converted value: 30

	num = 123
	str := strconv.Itoa(num)              // Convert int to string
	fmt.Println("Converted string:", str) // Output: Converted string: 123

	strNum := "456"
	numFromStr, err := strconv.Atoi(strNum) // Convert string to int
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("Converted int:", numFromStr) // Output: Converted int: 456
	}

	numfloat := 78.9
	numFromFloat := int(numfloat)                                 // Convert float64 to int
	fmt.Println("Converted int from float:", numFromFloat)        // Output: Converted int from float: 78
	numFromFloatStr := strconv.FormatFloat(numfloat, 'f', -1, 64) // Convert float64 to string
	fmt.Println("Converted string from float:", numFromFloatStr)  // Output: Converted string from float: 78.9
	strToFloat, err := strconv.ParseFloat(strNum, 64)             // Convert string to float64
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Println("Converted float from string:", strToFloat) // Output: Converted float from string: 456
	}
}

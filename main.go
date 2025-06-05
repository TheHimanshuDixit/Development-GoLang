package main

import (
	"fmt"
	"mylearning/myutil"
) // Importing the custom package
// main.go

func main() {
	myutil.PrintMessage() // Calling the function from the custom package

	var num int = 42
	fmt.Println("The number is:", num) // Using fmt package to print a message

	var str string = "Hello, World!"
	fmt.Println("The string is:", str) // Printing a string message

	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println("The array is:", arr) // Printing an array of integers

	var m = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("The map is:", m) // Printing a map with string keys and integer values

	var b bool = true
	fmt.Println("The boolean value is:", b) // Printing a boolean value

	var f float64 = 3.14
	fmt.Println("The float value is:", f) // Printing a float value

	var c complex128 = 1 + 2i
	fmt.Println("The complex number is:", c) // Printing a complex number

	var n = 20
	fmt.Println("The value of n is:", n) // Printing the value of n

	const pi = 3.14159
	fmt.Println("The value of pi is:", pi) // Printing a constant value

	v := 100
	fmt.Println("The value of v is:", v) // Printing the value of v

	// variable is capitial letter then it is exported else it is not exported
	var PublicVar = "This is a public variable"
	fmt.Println("Public variable:", PublicVar) // Printing a public variable
	var privateVar = "This is a private variable"
	fmt.Println("Private variable:", privateVar) // Printing a private variable
}

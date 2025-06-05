package main

import (
	"fmt"
)

// Defining a struct
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person_Details Person // Embedding the Person struct
	Position       string
	Salary         float64
}

func main() {

	// Creating an instance of the struct
	alice := Person{Name: "Alice", Age: 30}

	// Accessing struct fields
	fmt.Println("Name:", alice.Name)
	fmt.Println("Age:", alice.Age)

	var Bob = new(Person)
	Bob.Name = "Bob"
	Bob.Age = 25
	fmt.Println("Name:", Bob.Name)
	fmt.Println("Age:", Bob.Age)

	// Creating an instance of the Employee struct
	emp := Employee{
		Person_Details: Person{Name: "Charlie", Age: 28},
		Position:       "Developer",
		Salary:         75000,
	}
	fmt.Println("Employee Name:", emp.Person_Details.Name)
	fmt.Println("Employee Age:", emp.Person_Details.Age)
	fmt.Println("Employee Position:", emp.Position)
	fmt.Println("Employee Salary:", emp.Salary)
}

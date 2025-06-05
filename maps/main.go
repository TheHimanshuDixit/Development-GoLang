package main

import (
	"fmt"
)

func main() {
	// Creating a map
	ages := make(map[string]int)

	// Adding key-value pairs
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35

	// Accessing a value
	fmt.Println("Alice's age:", ages["Alice"])

	// Looping through the map
	for name, age := range ages {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	age, exists := ages["Bob"]
	if exists {
		fmt.Println("Bob's age:", age)
	} else {
		fmt.Println("Bob not found in the map")
	}

	// Deleting a key-value pair
	delete(ages, "Bob")
	fmt.Println("After deletion:", ages)
}

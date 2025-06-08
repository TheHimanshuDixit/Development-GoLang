package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	// Create a map to hold the JSON data
	// data := map[string]interface{}{
	// 	"name": "John Doe",
	// 	"age":  30,
	// 	"city": "New York",
	// }

	per := Person{
		Name: "John Doe",
		Age:  30,
		City: "New York",
	}

	// Convert the map to JSON
	jsonData, err := json.Marshal(per)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))

	// Unmarshal the JSON data back into a map
	var person Person
	err = json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the unmarshalled data
	fmt.Println("Unmarshalled Person struct:")
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("City:", person.City)
}

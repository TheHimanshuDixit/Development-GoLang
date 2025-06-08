package main

import (
	"encoding/json"
	"fmt"
	"strings"

	// "io/ioutil"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PerformGETRequest() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is OK (200)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}

	// Read the response body
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }

	// Print the response body as a string
	// fmt.Println(string(body))

	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// Print the Todo struct
	fmt.Printf("Todo: %+v\n", todo)
	// fmt.Printf("Todo ID: %d, Title: %s, Completed: %t\n", todo.ID, todo.Title, todo.Completed)
}

func PerformPOSTRequest() {
	// This function is a placeholder for future implementation
	fmt.Println("Performing POST request...")
	// You can implement the POST request logic here

	todo := Todo{
		UserID:    1,
		ID:        1,
		Title:     "Sample Todo",
		Completed: false,
	}

	// Convert the Todo struct to JSON
	todoJSON, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	jsonString := string(todoJSON)

	jsonReader := strings.NewReader(jsonString)

	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error performing POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is OK (200)
	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}

	// Read the response body
	var createdTodo Todo
	err = json.NewDecoder(resp.Body).Decode(&createdTodo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the created Todo struct
	fmt.Printf("Created Todo: %+v\n", createdTodo)
}

func PerformPUTRequest() {
	// This function is a placeholder for future implementation
	fmt.Println("Performing PUT request...")
	// You can implement the PUT request logic here

	todo := Todo{
		UserID:    23789,
		Title:     "Updated Todo",
		Completed: true,
	}

	// Convert the Todo struct to JSON
	todoJSON, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	jsonReader := strings.NewReader(string(todoJSON))

	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing PUT request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}

	var updatedTodo Todo
	err = json.NewDecoder(resp.Body).Decode(&updatedTodo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("Updated Todo: %+v\n", updatedTodo)
}

func PerformDELETERequest() {
	// This function is a placeholder for future implementation
	fmt.Println("Performing DELETE request...")
	// You can implement the DELETE request logic here

	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing DELETE request:", err)
		return
	}
	defer resp.Body.Close()

	// if resp.StatusCode != http.StatusNoContent {
	// 	fmt.Printf("Error: received status code %d\n", resp.StatusCode)
	// 	return
	// }

	fmt.Println("Todo deleted successfully")
}

func main() {
	// Call the PerformGETRequest function
	// PerformGETRequest()
	// Call the PerformPOSTRequest function
	// PerformPOSTRequest()
	// Call the PerformPUTRequest function
	// PerformPUTRequest()
	// Call the PerformDELETERequest function
	PerformDELETERequest()
}

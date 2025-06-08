package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
	time.Sleep(2 * time.Second) // Sleep to simulate work
	fmt.Println("Finished work in sayHello goroutine!")
}

func sayHi() {
	fmt.Println("Hi from goroutine!")
}

func main() {
	fmt.Println("Hello, World!")
	go sayHello()
	go sayHi()
	time.Sleep(3 * time.Second) // Wait for goroutines to finish
}

package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func(i int) {
			defer wg.Done() // Decrement the counter when the goroutine completes
			fmt.Printf("Goroutine %d is running\n", i)
		}(i)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines have completed.")
	fmt.Println("Exiting main function.")
}

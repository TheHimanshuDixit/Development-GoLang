package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name:")
	// var name string
	// fmt.Scanln(&name)
	// fmt.Printf("Hello, %s!\n", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s!\n", name)
}

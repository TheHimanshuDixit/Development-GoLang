package main

import (
	"fmt"
	"io/ioutil"
	// "io"
	// "os"
)

func main() {
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// defer file.Close()

	// // _, err := io.WriteString(file, "Hello, World!")
	// _, err = file.WriteString("Hello, World!")
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	return
	// }

	// fmt.Println("File created and written successfully.")

	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()

	// buffer := make([]byte, 1024)

	// for {
	// 	n, err := file.Read(buffer)
	// 	if err != nil {
	// 		if err.Error() == "EOF" {
	// 			break // End of file reached
	// 		}
	// 		fmt.Println("Error reading file:", err)
	// 		return
	// 	}
	// 	if n == 0 {
	// 		break // No more data to read
	// 	}
	// 	fmt.Print(string(buffer[:n]))
	// }

	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:")
	fmt.Println(string(content))

}

package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("https://example.com:8080/path?query=123#fragment")
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)
	fmt.Println("Query:", u.Query())
	fmt.Println("RawQuery:", u.RawQuery)
	fmt.Println("Fragment:", u.Fragment)

	u.Path = "/newpath"
	u.RawQuery = "newquery=456"

	newURL := u.String()
	fmt.Println("Updated URL:", newURL)
}

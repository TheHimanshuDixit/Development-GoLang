package main
import (
	"fmt"
	"strings"
	"unicode"
)

const goProgramming = "Go Programming"

func main() {
	fmt.Println(strings.ToUpper("hello"))	
	fmt.Println(strings.ToLower("WORLD"))
	fmt.Println(strings.TrimSpace("  Go Programming  "))
	fmt.Println(strings.Contains(goProgramming, "Go"))
	fmt.Println(strings.ReplaceAll("Go Go Go", "Go", "Golang"))
	fmt.Println(strings.Split("Go,Programming,Language", ","))
	fmt.Println(strings.Join([]string{"Go", "Programming", "Language"}, " "))
	fmt.Println(strings.Repeat("Go ", 3))
	fmt.Println(strings.Index(goProgramming, "Programming"))
	fmt.Println(strings.LastIndex(goProgramming, "o"))
	fmt.Println(strings.Count("Go Go Go", "Go"))
	fmt.Println(strings.HasPrefix(goProgramming, "Go"))
	fmt.Println(strings.HasSuffix(goProgramming, "Programming"))
	fmt.Println(strings.Fields("Go Programming Language"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "İSTANBUL"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "istanbul"))
}
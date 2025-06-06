package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	fmt.Println("Current Time (UTC):", currentTime.UTC())
	fmt.Println("Current Time (Location):", currentTime.In(time.Local))
	fmt.Println("Current Time (Formatted):", currentTime.Format("2006-01-02 15:04:05"))

	// 2006-01-02 15:04:05 is the reference time in Go, which is used for formatting.
	// It represents the date and time: January 2, 2006, at 15:04:05 (3:04:05 PM).
	// You can use this format to display the current time in a human-readable way.
	// Monday, January 2, 2006, 15:04:05 MST is the full format with timezone.
	fmt.Println("Current Time (Full Format):", currentTime.Format("Monday, January 2, 2006, 15:04:05 MST"))

	// time parse
	parsedTime, err := time.Parse("2006-01-02 15:04:05", "2023-10-01 12:00:00")
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed Time:", parsedTime)

	// add duration
	addedTime := currentTime.Add(2 * time.Hour)
	fmt.Println("Time after adding 2 hours:", addedTime)
	// subtract duration	
	subtractedTime := currentTime.Add(-30 * time.Minute)
	fmt.Println("Time after subtracting 30 minutes:", subtractedTime)
	// sleep for 2 seconds
	time.Sleep(2 * time.Second)
	fmt.Println("Slept for 2 seconds, current time:", time.Now())
	// sleep for 2 seconds with a ticker
	ticker := time.NewTicker(2 * time.Second)
	for t := range ticker.C {
		fmt.Println("Ticker ticked at:", t)
		if t.After(currentTime.Add(10*time.Second)) {
			ticker.Stop()
			break
		}
	}
}

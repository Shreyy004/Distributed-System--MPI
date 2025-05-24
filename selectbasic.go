package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine to send data after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Data from ch1"
	}()

	// Goroutine to send data after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Data from ch2"
	}()

	// Use select to wait for multiple channels
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}
}

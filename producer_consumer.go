package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Produced:", i)
		ch <- i // Send data to channel (blocks if full)
		time.Sleep(time.Second) // Simulate production time
	}
	close(ch) // Close channel to signal no more data
}

func consumer(ch chan int, done chan bool) {
	for item := range ch { // Receive data (blocks if empty)
		fmt.Println("Consumed:", item)
		time.Sleep(2 * time.Second) // Simulate processing time
	}
	done <- true // Signal completion
}

func main() {
	bufferSize := 2
	ch := make(chan int, bufferSize) // Buffered channel (queue size = 2)
	done := make(chan bool)

	go producer(ch)
	go consumer(ch, done)

	<-done // Wait for consumer to finish
	fmt.Println("All items processed!")
}

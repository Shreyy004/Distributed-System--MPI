package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan string) {
	fmt.Printf("Worker %d: started\n", id)
	time.Sleep(time.Second)                      // Simulate some work
	ch <- fmt.Sprintf("Worker %d: finished", id) // Send message to channel
}

func main() {
	ch := make(chan string) // Create an unbuffered channel

	for i := 1; i <= 3; i++ {
		go worker(i, ch) // Start goroutine
	}

	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch) // Receive messages from the channel
	}

	fmt.Println("All workers finished.")
}

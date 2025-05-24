package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan string) {
	fmt.Printf("Worker %d: started\n", id)
	time.Sleep(time.Second)
	ch <- fmt.Sprintf("Worker %d: finished", id) // Send message to channel
}

func main() {
	ch := make(chan string, 3) // Buffered channel with capacity 3

	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	time.Sleep(2 * time.Second) // Wait for workers
	close(ch)                   // Close the channel

	for msg := range ch { // Receive values until the channel is closed
		fmt.Println(msg)
	}
}

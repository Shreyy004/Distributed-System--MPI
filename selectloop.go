package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			ch1 <- "Message from ch1"
		}
	}()

	go func() {
		for {
			time.Sleep(3 * time.Second)
			ch2 <- "Message from ch2"
		}
	}()

	// Continuously listen to both channels
	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}

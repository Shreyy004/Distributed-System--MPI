package main

import (
	"fmt"
	"sync"
)

// Order struct to store order details
type Order struct {
	id   int
	user string
	item string
}

// Function to process orders
func processOrders(orderChannel chan Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderChannel {
		fmt.Printf("Processing Order #%d: %s ordered %s\n", order.id, order.user, order.item)
	}
}

func main() {
	// Buffered channel to hold orders
	orderChannel := make(chan Order, 3)

	var wg sync.WaitGroup

	// Sample orders
	orders := []Order{
		{1, "Alice", "Pizza"},
		{2, "Bob", "Burger"},
		{3, "Charlie", "Pasta"},
	}

	// Start processing orders in a separate goroutine
	wg.Add(1)
	go processOrders(orderChannel, &wg)

	// Place orders
	for _, order := range orders {
		fmt.Printf("%s placed an order for %s\n", order.user, order.item)
		orderChannel <- order
	}

	// Close the channel and wait for processing to finish
	close(orderChannel)
	wg.Wait()
}

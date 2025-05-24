package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate microservices
func checkInventory(orderID int) string {
	time.Sleep(100 * time.Millisecond) // Simulate network latency
	if orderID%2 == 0 {
		return "" // Inventory available
	}
	return "inventory not available"
}

func processOrder(orderID int) string {
	time.Sleep(150 * time.Millisecond) // Simulate processing time
	return ""                          // No error
}

func validatePayment(orderID int) string {
	time.Sleep(200 * time.Millisecond) // Simulate payment validation
	if orderID%3 == 0 {
		return "payment validation failed"
	}
	return "" // No error
}

func placeOrder(orderID int) []string {
	var wg sync.WaitGroup
	errChan := make(chan string, 2) // Buffer for 2 errors at max

	// Concurrently check inventory
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := checkInventory(orderID); err != "" {
			errChan <- fmt.Sprintf("inventory check failed: %s", err)
		}
	}()

	// Concurrently validate payment
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := validatePayment(orderID); err != "" {
			errChan <- fmt.Sprintf("payment validation failed: %s", err)
		}
	}()
	wg.Wait()
	close(errChan)

	// Collect errors
	var errors []string
	for err := range errChan {
		errors = append(errors, err)
	}

	// If errors exist, return them
	if len(errors) > 0 {
		return errors
	}

	// If no errors, process the order
	if err := processOrder(orderID); err != "" {
		return []string{fmt.Sprintf("order processing failed: %s", err)}
	}

	return nil // No error, order placed successfully
}

func main() {
	orderIDs := []int{1, 2, 3, 4, 5, 6}

	for _, orderID := range orderIDs {
		fmt.Printf("Placing order %d...\n", orderID)
		errors := placeOrder(orderID)
		if errors != nil {
			for _, err := range errors {
				fmt.Printf("Order %d failed: %v\n", orderID, err)
			}
		} else {
			fmt.Printf("Order %d placed successfully!\n", orderID)
		}
	}
}

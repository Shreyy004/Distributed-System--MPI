// Consider a pizza-making process in a restaurant where multiple workers handle different stages of preparation. Each stage must be completed before moving to the next one. The process consists of the following steps in order:

// Dough Preparation
// Sauce and Toppings Application
// Baking the Pizza
// Packing for Delivery
// Implement a Go program using goroutines and synchronization techniques to ensure that each step starts only after the previous one has been completed.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Step 1: Dough Preparation
func prepareDough(wg *sync.WaitGroup) {
	defer wg.Done() // Mark this step as done when it completes
	fmt.Println("Dough Preparation started...")
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Println("Dough Preparation completed!")
}

// Step 2: Sauce and Toppings Application (Waits for Dough)
func applySauceAndToppings(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Applying Sauce and Toppings...")
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Println("Sauce and Toppings applied!")
}

// Step 3: Baking the Pizza (Waits for Sauce & Toppings)
func bakePizza(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Baking the Pizza...")
	time.Sleep(3 * time.Second) // Simulate work
	fmt.Println("Pizza Baked!")
}

// Step 4: Packing for Delivery (Waits for Baking)
func packForDelivery(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Packing for Delivery...")
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Println("Pizza Packed and Ready for Delivery!")
}

func main() {
	var wg sync.WaitGroup

	// Run each step sequentially
	wg.Add(1)
	go prepareDough(&wg)
	wg.Wait() // Wait for Dough Preparation to finish

	wg.Add(1)
	go applySauceAndToppings(&wg)
	wg.Wait() // Wait for Sauce & Toppings to finish

	wg.Add(1)
	go bakePizza(&wg)
	wg.Wait() // Wait for Baking to finish

	wg.Add(1)
	go packForDelivery(&wg)
	wg.Wait() // Wait for Packing to finish

	fmt.Println("Pizza-making process completed successfully!")
}

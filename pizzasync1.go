// Consider a pizza-making process in a restaurant where multiple workers handle different stages of preparation. Each stage must be completed before moving to the next one. The process consists of the following steps in order:

// Dough Preparation
// Sauce and Toppings Application
// Baking the Pizza
// Packing for Delivery
// Implement a Go program using goroutines and synchronization techniques to ensure that each step starts only after the previous one has been completed.

package main

import (
	"fmt"
	"time"
)

// Step 1: Dough Preparation
func prepareDough(done chan bool) {
	fmt.Println("Dough Preparation started...")
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Println("Dough Preparation completed!")
	done <- true // Signal completion
}

// Step 2: Sauce and Toppings Application (Waits for Dough)
func applySauceAndToppings(prevDone chan bool, done chan bool) {
	<-prevDone // Wait for Dough Preparation to finish
	fmt.Println("Applying Sauce and Toppings...")
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Println("Sauce and Toppings applied!")
	done <- true // Signal completion
}

// Step 3: Baking the Pizza (Waits for Sauce & Toppings)
func bakePizza(prevDone chan bool, done chan bool) {
	<-prevDone // Wait for Sauce & Toppings to finish
	fmt.Println("Baking the Pizza...")
	time.Sleep(3 * time.Second) // Simulate work
	fmt.Println("Pizza Baked!")
	done <- true // Signal completion
}

// Step 4: Packing for Delivery (Waits for Baking)
func packForDelivery(prevDone chan bool, done chan bool) {
	<-prevDone // Wait for Baking to finish
	fmt.Println("Packing for Delivery...")
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Println("Pizza Packed and Ready for Delivery!")
	done <- true // Signal completion
}

func main() {
	// Creating channels for step-wise synchronization
	doughDone := make(chan bool)
	sauceDone := make(chan bool)
	bakingDone := make(chan bool)
	packingDone := make(chan bool)

	// Running each step in a goroutine
	go prepareDough(doughDone)
	go applySauceAndToppings(doughDone, sauceDone)
	go bakePizza(sauceDone, bakingDone)
	go packForDelivery(bakingDone, packingDone)

	// Wait for the final step (Packing) to finish
	<-packingDone

	fmt.Println("Pizza-making process completed successfully!")
}

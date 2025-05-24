package main

import (
	"fmt"
	"sync"
)

// Struct to hold construction type and its rate per square foot
type ConstructionType struct {
	name string
	rate float64
}

// Mutex to ensure safe concurrent writes
var mu sync.Mutex

// Function to calculate budget
func calculateBudget(area float64, construction ConstructionType, wg *sync.WaitGroup) {
	defer wg.Done()
	budget := area * construction.rate

	// Locking to ensure only one goroutine prints at a time
	mu.Lock()
	fmt.Printf("%s: $%.2f\n", construction.name, budget)
	mu.Unlock()
}

func printHouse(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Calculating budget for House (Villa)...")
}

func printFlats(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Calculating budget for Flats...")
}

func printCommercial(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Calculating budget for Commercial Buildings...")
}

func main() {
	var choice int
	var area float64

	fmt.Print("Enter the area in square feet: ")
	fmt.Scanln(&area)

	fmt.Println("Choose the type of construction project:")
	fmt.Println("1. House (Villa)")
	fmt.Println("2. Flats")
	fmt.Println("3. Commercial Buildings")
	fmt.Println("4. All in parallel")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	// Define construction types
	house := ConstructionType{"House (Villa)", 100.0}
	flats := ConstructionType{"Flats", 120.0}
	commercial := ConstructionType{"Commercial Buildings", 200.0}

	var wg sync.WaitGroup

	switch choice {
	case 1:
		wg.Add(1)
		go calculateBudget(area, house, &wg)
	case 2:
		wg.Add(1)
		go calculateBudget(area, flats, &wg)
	case 3:
		wg.Add(1)
		go calculateBudget(area, commercial, &wg)
	case 4:
		wg.Add(3)
		go printHouse(&wg)
		go printFlats(&wg)
		go printCommercial(&wg)
		wg.Wait() // Ensure printing is completed before calculations start

		wg.Add(3)
		go calculateBudget(area, house, &wg)
		go calculateBudget(area, flats, &wg)
		go calculateBudget(area, commercial, &wg)
	default:
		fmt.Println("Invalid choice")
		return
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("Budget estimation completed.")
}

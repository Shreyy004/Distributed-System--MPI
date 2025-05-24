package main

import (
	"fmt"
	"sync"
	"time"
)

type ConstructionType struct {
	name string
	rate float64
}

var mu sync.Mutex

func calculateBudget(area float64, bud ConstructionType, wg *sync.WaitGroup) {
	defer wg.Done()
	budget := area * bud.rate

	mu.Lock()
	fmt.Printf("%s: $%.2f\n", bud.name, budget)
	mu.Unlock()

}

func housebud(wg *sync.WaitGroup) {

	defer wg.Done()
	time.Sleep(time.Microsecond)
	fmt.Println("inside house")

}

func flatbud(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Microsecond)
	fmt.Println("inside flat")
}

func buildbud(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Microsecond)
	fmt.Println("inside commercial building")
}

func main() {

	var wg sync.WaitGroup
	var choice int
	var area float64

	fmt.Print("enter the area: ")
	fmt.Scanln(&area)

	fmt.Println("Choose the type of construction project:")
	fmt.Println("1. House (Villa)")
	fmt.Println("2. Flats")
	fmt.Println("3. Commercial Buildings")
	fmt.Println("4. All in parallel")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	house := ConstructionType{"(house)villa", 100.0}
	flat := ConstructionType{"flat", 200.0}
	com := ConstructionType{"commercial", 250.0}

	switch choice {
	case 1:
		wg.Add(1)
		go housebud(&wg)
	case 2:
		wg.Add(1)
		go flatbud(&wg)
	case 3:
		wg.Add(1)
		go buildbud(&wg)
	case 4:
		wg.Add(3)
		go housebud(&wg)
		go flatbud(&wg)
		go buildbud(&wg)

		wg.Wait()

		wg.Add(3)
		go calculateBudget(area, house, &wg)
		go calculateBudget(area, flat, &wg)
		go calculateBudget(area, com, &wg)
	default:
		fmt.Println("invalid choice")
		return

	}
	wg.Wait()
	fmt.Println("budget calculation done")
}

package main

import (
	"fmt"
	"sync"
	"time"
)

// Total seats available
var availableSeats int = 10

// Mutex for safe seat booking
var mu sync.Mutex

// Function to book tickets
func bookTicket(user string, tickets int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate network delay
	time.Sleep(time.Millisecond * 200)

	// Lock before checking and updating available seats
	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("%s is trying to book %d ticket(s)...\n", user, tickets)

	// Check seat availability
	if tickets <= availableSeats {
		availableSeats -= tickets
		fmt.Printf("✅ Booking confirmed for %s! %d seat(s) booked. Remaining seats: %d\n", user, tickets, availableSeats)
	} else {
		fmt.Printf("❌ Booking failed for %s! Not enough seats available.\n", user)
	}
}

func main() {
	var wg sync.WaitGroup

	// Simulating multiple users booking tickets concurrently
	users := []struct {
		name    string
		tickets int
	}{
		{"Alice", 3},
		{"Bob", 5},
		{"Charlie", 2},
		{"David", 4},
		{"Emma", 2},
	}

	// Booking tickets concurrently
	for _, user := range users {
		wg.Add(1)
		go bookTicket(user.name, user.tickets, &wg)
	}

	// Wait for all bookings to finish
	wg.Wait()

	fmt.Println(" All bookings processed.")
}

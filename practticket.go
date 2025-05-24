package main

import (
	"fmt"
	"sync"
	"time"
)

type bookings struct {
	availableSeats int
	mu             sync.Mutex
}

func (b *bookings) bookTicket(user string, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	b.mu.Lock()
	defer b.mu.Unlock()
	if n > b.availableSeats {
		fmt.Println("tickets full")
		return
	}
	b.availableSeats -= n
	fmt.Printf("%s booked %d tickets\n", user, n)
	time.Sleep(time.Millisecond)
}

func main() {
	av := bookings{availableSeats: 10}
	var wg sync.WaitGroup
	users := []struct {
		name string
		tk   int
	}{
		{"first", 3},
		{"second", 2},
		{"third", 3},
		{"four", 2},
		{"five", 2},
	}
	for _, user := range users {
		wg.Add(1)
		go av.bookTicket(user.name, user.tk, &wg)
	}
	wg.Wait()
	fmt.Println("booked finished")

}

package main

import (
	"fmt"
	"sync"
	"time"
)

type Safecounter struct {
	mu  sync.Mutex
	val int
}

func (c *Safecounter) Increment() {
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
}

func (c *Safecounter) getval() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}
func main() {

	var wg sync.WaitGroup
	counter := Safecounter{}

	for i := 1; i <= 5; i++ {

		wg.Add(1)
		go func() {

			defer wg.Done()
			for j := 1; j <= 10; j++ {
				counter.Increment()
				time.Sleep(time.Millisecond)
			}
		}()
	}
	wg.Wait()
	fmt.Println("final counter value:", counter.getval())

}

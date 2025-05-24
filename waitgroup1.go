package main

import (
	"fmt"
	"sync"
	"time"
)

func workers(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(name, " worker started")
	time.Sleep(time.Millisecond)
	fmt.Println(name, "completed")
}

func main() {
	var wg sync.WaitGroup
	var tasks = []string{"taskA", "taskB", "taskC"}

	for _, t := range tasks {
		wg.Add(1)
		go workers(t, &wg)

	}
	wg.Wait()
	fmt.Println("All tasks finished.")
}

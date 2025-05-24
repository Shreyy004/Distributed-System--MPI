package main

import (
      "fmt"
      "time"
)

func main() {
    tick := time.Tick(100 * time.Millisecond) //returns a channel sends time every 100 milliseconds
    boom := time.After(500 * time.Millisecond) //returns a channel , sends a single value after 500 ms and then closes 
    
    for {
      select {
        case <- tick:
             fmt.Println("tick. ")
        case <- boom:
             fmt.Println("BOOM!")
             return
        default:
             fmt.Println("   .")
             time.Sleep(50 * time.Millisecond)
      }
    }
}



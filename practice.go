package main

import (
    "fmt"
    "sync"
)

func assembleStep(abc string, x chan bool, y chan bool, w *sync.WaitGroup) {
    defer w.Done()
    if x != nil {
        val := <-x
        fmt.Printf("received: %v\n", val)
    }
    fmt.Println(abc, "completed")
    if y != nil {
        y <- true // Send completion signal to next step
    }
}

func main() {
    var wg sync.WaitGroup

    engineDone := make(chan bool)
    abcDone := make(chan bool)
    bodyDone := make(chan bool)
    steeringDone := make(chan bool)

    wg.Add(4)

    go func() { engineDone <- true }() // Start the first step

    go assembleStep("engine setup", engineDone, abcDone, &wg)
    go assembleStep("accelerator-brake-clutch setup", abcDone, bodyDone, &wg)
    go assembleStep("body setup", bodyDone, steeringDone, &wg)
    go assembleStep("steering setup", steeringDone, nil, &wg)

    wg.Wait()
    fmt.Println("2 wheeler assembly complete")
}

package main
import (
"fmt"
"time"
)

func engine_setUp(Done chan bool) {
	fmt.Println("setting up engine...")
	time.Sleep(1*time.Second)
	fmt.Println("engine set up complete!")
	Done <- true
}

func ABC_setUp(engineDone chan bool, abcDone chan bool) {
	<-engineDone // Wait for engine setup
	fmt.Println("setting up accelerator-brake-clutch...")
	time.Sleep(1*time.Second)
	fmt.Println("accelerator-brake-clutch set up complete!")
	abcDone <- true
}

func assemble() {
	engineDone := make(chan bool)
	abcDone := make(chan bool)

	go engine_setUp(engineDone)
	go ABC_setUp(engineDone, abcDone)

	<-abcDone // Wait for ABC setup to finish
	fmt.Println("complete")
}


func main() {
	var looper int
	fmt.Println("if you wish to order bike, enter 1; else -1")
	fmt.Scanf("%d", &looper)
	if looper == 1 {
	assemble()
	} else {
	fmt.Println("exiting...")
	}
}

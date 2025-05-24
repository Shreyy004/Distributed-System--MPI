package main

import (
	"fmt"
	"net/rpc"
	"sync"
)

// Request struct for booking or canceling a car
type Request struct {
	Category string
}

// Response struct to receive responses from the server
type Response struct {
	Message string
}

func main() {
	// Connect to RPC server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	// WaitGroup to manage Goroutines
	var wg sync.WaitGroup

	// Function to check availability
	wg.Add(1)
	go func() {
		defer wg.Done()
		var res Response
		err := client.Call("RentalService.CheckAvailability", struct{}{}, &res)
		if err != nil {
			fmt.Println("Availability Error:", err)
		} else {
			fmt.Println("Availability:\n", res.Message)
		}
	}()

	// Function to book a car
	wg.Add(1)
	go func() {
		defer wg.Done()
		var res Response
		req := Request{Category: "SUV"}
		err := client.Call("RentalService.BookCar", req, &res)
		if err != nil {
			fmt.Println("Booking Error:", err)
		} else {
			fmt.Println("Booking Response:", res.Message)
		}
	}()

	// Function to cancel booking
	wg.Add(1)
	go func() {
		defer wg.Done()
		var res Response
		req := Request{Category: "SUV"}
		err := client.Call("RentalService.CancelBooking", req, &res)
		if err != nil {
			fmt.Println("Cancellation Error:", err)
		} else {
			fmt.Println("Cancellation Response:", res.Message)
		}
	}()

	// Wait for all Goroutines to finish
	wg.Wait()
}

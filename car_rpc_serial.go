package main

import (
	"fmt"
	"net/rpc"
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

	// Step 1: Check availability
	var res Response
	err = client.Call("RentalService.CheckAvailability", struct{}{}, &res)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Availability:\n", res.Message)
	}

	// Step 2: Book a car
	req := Request{Category: "SUV"}
	err = client.Call("RentalService.BookCar", req, &res)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Booking Response:", res.Message)
	}

	// Step 3: Cancel the booking
	err = client.Call("RentalService.CancelBooking", req, &res)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cancellation Response:", res.Message)
	}
}

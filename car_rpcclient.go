package main

import (
	"fmt"
	"net/rpc"
)

// Request struct for booking or canceling a car
type Request struct {
	Category string
}

// Response struct to receive responses from server
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

	for {
		fmt.Println("\nCar Rental System:")
		fmt.Println("1. Book a Car")
		fmt.Println("2. Cancel a Booking")
		fmt.Println("3. Check Availability")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter car category to book: ")
			var category string
			fmt.Scanln(&category)

			req := Request{Category: category}
			var res Response
			err = client.Call("RentalService.BookCar", req, &res)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(res.Message)
			}

		case 2:
			fmt.Print("Enter car category to cancel booking: ")
			var category string
			fmt.Scanln(&category)

			req := Request{Category: category}
			var res Response
			err = client.Call("RentalService.CancelBooking", req, &res)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(res.Message)
			}

		case 3:
			var res Response
			err = client.Call("RentalService.CheckAvailability", struct{}{}, &res)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(res.Message)
			}

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice! Try again.")
		}
	}
}

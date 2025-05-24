package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

// Car represents a car category with available units and rental price
type Car struct {
	Available int
	Price     int
}

// RentalService holds the car inventory
type RentalService struct {
	Inventory map[string]*Car
	mu        sync.Mutex
}

// Request struct for booking or canceling a car
type Request struct {
	Category string
}

// Response struct to send availability or booking confirmation
type Response struct {
	Message string
}

// Book a car if available
func (r *RentalService) BookCar(req Request, res *Response) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if car, exists := r.Inventory[req.Category]; exists {
		if car.Available > 0 {
			car.Available--
			res.Message = fmt.Sprintf("Car booked successfully in %s category!", req.Category)
		} else {
			res.Message = "No cars available in this category!"
		}
	} else {
		res.Message = "Invalid car category!"
	}
	return nil
}

// Cancel a booked car
func (r *RentalService) CancelBooking(req Request, res *Response) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if car, exists := r.Inventory[req.Category]; exists {
		car.Available++
		res.Message = fmt.Sprintf("Booking canceled for %s category!", req.Category)
	} else {
		res.Message = "Invalid car category!"
	}
	return nil
}

// Check car availability
func (r *RentalService) CheckAvailability(_ struct{}, res *Response) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	message := "Car Availability:\n"
	for category, car := range r.Inventory {
		message += fmt.Sprintf("%s: %d available at Rs.%d per day\n", category, car.Available, car.Price)
	}
	res.Message = message
	return nil
}

func main() {
	// Create a new rental service
	service := &RentalService{
		Inventory: map[string]*Car{
			"Economy":      {10, 800},
			"Compact":      {15, 1000},
			"Intermediate": {8, 1500},
			"Full-Size":    {5, 2000},
			"SUV":          {3, 2500},
		},
	}

	// Register the service
	rpc.Register(service)

	// Start listening for client requests
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Car Rental RPC Server is running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

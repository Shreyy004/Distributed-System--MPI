package main

import (
	"fmt"
	"net/rpc"
)

// Define arguments struct
type Args struct {
	A, B int
}

func main() {
	// Connect to the RPC server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to RPC server:", err)
		return
	}

	// Define input arguments
	args := Args{A: 6, B: 7}
	var result int

	// Call the Multiply function
	err = client.Call("Arithmetic.Multiply", &args, &result)
	if err != nil {
		fmt.Println("Error calling RPC method:", err)
		return
	}

	fmt.Println("Multiplication Result:", result)
}

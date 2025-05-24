package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to RPC server:", err)
		return
	}
	args := Args{A: 5, B: 10}
	var res int

	err = client.Call("Calculator.Add", &args, &res)
	if err != nil {
		fmt.Println("failed to calculate")
	}
	fmt.Println("addition res: ", res)

	err = client.Call("Calculator.Sub", &args, &res)
	if err != nil {
		fmt.Println("failed to calculate")
	}
	fmt.Println("sub res: ", res)

}

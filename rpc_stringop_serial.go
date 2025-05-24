package main

import (
	"fmt"
	"net/rpc"
)

// CallService calls an RPC method
func CallService(client *rpc.Client, method string, request StringRequest) {
	var response StringResponse
	err := client.Call("StringService."+method, request, &response)
	if err != nil {
		fmt.Println("Error calling", method, ":", err)
		return
	}
	fmt.Println("Result of", method, ":", response.Result)
}

type StringRequest struct {
	Str   string
	Char  string
	Index int
	Str2  string
}

type StringResponse struct {
	Result string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	// Calls in serial order
	CallService(client, "Insert", StringRequest{Str: "hello", Char: "X", Index: 2})
	CallService(client, "Delete", StringRequest{Str: "hello", Index: 1})
	CallService(client, "Modify", StringRequest{Str: "hello", Char: "Y", Index: 3})
	CallService(client, "Concat", StringRequest{Str: "hello", Str2: "world"})
}

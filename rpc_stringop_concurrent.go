package main

import (
	"fmt"
	"net/rpc"
	"sync"
)

func CallService(client *rpc.Client, method string, request StringRequest, wg *sync.WaitGroup) {
	defer wg.Done()

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

	var wg sync.WaitGroup
	wg.Add(4)

	// Concurrently calling the services
	go CallService(client, "Insert", StringRequest{Str: "hello", Char: "X", Index: 2}, &wg)
	go CallService(client, "Delete", StringRequest{Str: "hello", Index: 1}, &wg)
	go CallService(client, "Modify", StringRequest{Str: "hello", Char: "Y", Index: 3}, &wg)
	go CallService(client, "Concat", StringRequest{Str: "hello", Str2: "world"}, &wg)

	wg.Wait()
}

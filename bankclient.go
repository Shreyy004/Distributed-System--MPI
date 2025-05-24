package main

import (
	"fmt"
	"net/rpc"
	"sync"
)

// Structs matching the server's arguments and response
type TransactionArgs struct {
	Amount int
}

type BalanceResponse struct {
	Balance int
}

// Function to make RPC requests
func makeRequest(client *rpc.Client, method string, args interface{}, reply interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	err := client.Call(method, args, reply)
	if err != nil {
		fmt.Println("Error calling RPC method:", err)
	} else {
		fmt.Printf("%s Response: %+v\n", method, reply)
	}
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to RPC server:", err)
		return
	}
	defer client.Close()

	var wg sync.WaitGroup

	// Create multiple requests
	transactions := []struct {
		method string
		args   interface{}
		reply  interface{}
	}{
		{"BankAccount.Deposit", &TransactionArgs{Amount: 500}, &BalanceResponse{}},
		{"BankAccount.Withdraw", &TransactionArgs{Amount: 200}, &BalanceResponse{}},
		{"BankAccount.CheckBalance", struct{}{}, &BalanceResponse{}},
		{"BankAccount.Deposit", &TransactionArgs{Amount: 300}, &BalanceResponse{}},  // Exceeds limit
		{"BankAccount.Withdraw", &TransactionArgs{Amount: 100}, &BalanceResponse{}}, // Exceeds limit
	}

	// Make requests concurrently
	for _, tx := range transactions {
		wg.Add(1)
		go makeRequest(client, tx.method, tx.args, tx.reply, &wg)
	}

	wg.Wait()
}

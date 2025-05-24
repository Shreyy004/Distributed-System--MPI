package main

import (
	"fmt"
	"net/rpc"
	"sync"
)

type TransactionArgs struct {
	Amount int
}
type BalanceResponse struct {
	Balance int
}

func makeRequest(client *rpc.Client, method string, args interface{}, reply interface{}, wg *sync.WaitGroup) {

}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("cant connect to client")
		return
	}
	defer client.Close()
	var wg sync.WaitGroup
	transactions := []struct {
		method string
		args   interface{}
		reply  interface{}
	}{
		{"BankAccount.Deposit", &TransactionArgs{Amount: 500}, &BalanceResponse{}},
		{"BankAccount.Withdraw", &TransactionArgs{Amount: 100}, &BalanceResponse{}},
		{"BankAccount.Deposit", &TransactionArgs{Amount: 500}, &BalanceResponse{}},
		{"BankAccount.CheckBalance", struct{}{}, &BalanceResponse{}},
		{"BankAccount.Deposit", &TransactionArgs{Amount: 300}, &BalanceResponse{}},  // Exceeds limit
		{"BankAccount.Withdraw", &TransactionArgs{Amount: 100}, &BalanceResponse{}}, // Exceeds limit

	}
	for _, tx := range transactions {
		wg.Add(1)
		go makeRequest(client, tx.method, tx.args, tx.reply, &wg)
	}
	wg.Wait()
}

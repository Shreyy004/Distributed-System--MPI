package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

// BankAccount represents a shared bank account
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

// Arguments structure for deposit and withdrawal
type TransactionArgs struct {
	Amount int
}

// Response structure for returning balance
type BalanceResponse struct {
	Balance int
}

// Semaphore to limit concurrent requests (allow max 3 at a time)
var sem = make(chan struct{}, 3) // Allows only 3 concurrent requests

// Deposit method
func (b *BankAccount) Deposit(args *TransactionArgs, reply *BalanceResponse) error {
	sem <- struct{}{}        // Acquire slot
	defer func() { <-sem }() // Release slot

	b.mu.Lock() // Lock to ensure consistency
	defer b.mu.Unlock()

	b.balance += args.Amount
	reply.Balance = b.balance
	fmt.Println("Deposited:", args.Amount, "New Balance:", b.balance)
	return nil
}

// Withdraw method
func (b *BankAccount) Withdraw(args *TransactionArgs, reply *BalanceResponse) error {
	sem <- struct{}{}
	defer func() { <-sem }()

	b.mu.Lock()
	defer b.mu.Unlock()

	if args.Amount > b.balance {
		return errors.New("insufficient funds")
	}

	b.balance -= args.Amount
	reply.Balance = b.balance
	fmt.Println("Withdrawn:", args.Amount, "New Balance:", b.balance)
	return nil
}

// CheckBalance method
func (b *BankAccount) CheckBalance(_ struct{}, reply *BalanceResponse) error {
	sem <- struct{}{}
	defer func() { <-sem }()

	b.mu.Lock()
	defer b.mu.Unlock()

	reply.Balance = b.balance
	fmt.Println("Checked Balance:", b.balance)
	return nil
}

func main() {
	account := new(BankAccount)
	account.balance = 1000 // Initial balance

	server := rpc.NewServer()
	err := server.Register(account)
	if err != nil {
		fmt.Println("Error registering RPC service:", err)
		return
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Banking RPC Server running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		go server.ServeConn(conn)
	}
}

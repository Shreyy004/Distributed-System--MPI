package main

import (
	"fmt"
	"sync"
	"time"
)

// Struct for Bank Account
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

// Deposit function
func (acc *BankAccount) deposit(user string, amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	acc.mu.Lock()
	defer acc.mu.Unlock()

	// Process deposit
	fmt.Printf(" %s deposited $%d\n", user, amount)
	acc.balance += amount

	// Simulating processing time
	time.Sleep(500 * time.Millisecond)
}

// Withdraw function
func (acc *BankAccount) withdraw(user string, amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	acc.mu.Lock()
	defer acc.mu.Unlock()

	// Check if withdrawal is possible
	if amount > acc.balance {
		fmt.Printf(" %s failed to withdraw $%d (Insufficient balance!)\n", user, amount)
		return
	}

	// Process withdrawal
	fmt.Printf(" %s withdrew $%d\n", user, amount)
	acc.balance -= amount

	// Simulating processing time
	time.Sleep(500 * time.Millisecond)
}

func main() {
	// Initialize bank account with $1000
	account := BankAccount{balance: 1000}
	var wg sync.WaitGroup

	// Simulating multiple users making transactions
	users := []struct {
		name   string
		action string
		amount int
	}{
		{"Alice", "deposit", 500},
		{"Bob", "withdraw", 300},
		{"Charlie", "withdraw", 800}, // This may fail
		{"David", "deposit", 200},
		{"Emma", "withdraw", 400},
	}

	// Start goroutines for transactions
	for _, user := range users {
		wg.Add(1)
		if user.action == "deposit" {
			go account.deposit(user.name, user.amount, &wg)
		} else {
			go account.withdraw(user.name, user.amount, &wg)
		}
	}

	// Wait for all transactions to finish
	wg.Wait()

	// Final account balance
	fmt.Printf("\n✅ Final Account Balance: $%d\n", account.balance)
}

package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type bankdetails struct {
	balance int
	mu      sync.Mutex
}

type TransactionArgs struct {
	Amount int
}
type BalanceResponse struct {
	Balance int
}

var sem = make(chan struct{}, 3)

func (b *bankdetails) Deposit(args *TransactionArgs, reply *BalanceResponse) error {
	sem <- struct{}{}
	defer func() {
		<-sem
	}()

	b.balance += args.Amount
	reply.Balance = b.balance
	fmt.Println("Deposited:", args.Amount, "New Balance:", b.balance)
	return nil
}

func (b *bankdetails) Withdraw(args *TransactionArgs, reply *BalanceResponse) error {
	sem <- struct{}{}
	defer func() {
		<-sem
	}()
	if args.Amount > b.balance {
		return errors.New("insuffiencient balance")
	}

	b.balance -= args.Amount
	reply.Balance = b.balance
	fmt.Println("withdrew:", args.Amount, "New Balance:", b.balance)
	return nil
}

func (b *bankdetails) CheckBalance(_ struct{}, reply *BalanceResponse) error {

	sem <- struct{}{}
	defer func() {
		<-sem
	}()
	reply.Balance = b.balance
	fmt.Println("balance: ", b.balance)
	return nil
}

func main() {

	account := new(bankdetails)
	account.balance = 1000
	Server := rpc.NewServer()
	Server.Register(account)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("server not found")
	}
	defer listener.Close()
	fmt.Println("server connected to port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go rpc.ServeConn(conn)

	}

}

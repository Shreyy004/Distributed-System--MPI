package main

import (
	"fmt"
	"sync"
	"time"
)

// Define Order struct
type Order struct {
	OrderID     int
	OrderType   string  // "buy" or "sell"
	StockSymbol string  // Stock ticker like "AAPL"
	Quantity    int     // Number of stocks
	Price       float64 // Price per stock
}

var buyOrders []Order
var sellOrders []Order
var mu sync.Mutex

// AddOrder function using struct instead of interface{}
func AddOrder(order Order) {
	mu.Lock()
	defer mu.Unlock()

	if order.OrderType == "buy" {
		buyOrders = append(buyOrders, order)
		fmt.Printf("Added buy order: %+v\n", order)
	} else if order.OrderType == "sell" {
		sellOrders = append(sellOrders, order)
		fmt.Printf("Added sell order: %+v\n", order)
	}
}

// MatchOrders function using struct
func MatchOrders() {
	mu.Lock()
	defer mu.Unlock()

	for i := 0; i < len(buyOrders); i++ {
		for j := 0; j < len(sellOrders); j++ {
			if buyOrders[i].StockSymbol == sellOrders[j].StockSymbol &&
				buyOrders[i].Price >= sellOrders[j].Price &&
				buyOrders[i].Quantity == sellOrders[j].Quantity {

				fmt.Printf("Trade executed: Buy Order %v and Sell Order %v\n", buyOrders[i].OrderID, sellOrders[j].OrderID)

				// Remove matched orders
				buyOrders = append(buyOrders[:i], buyOrders[i+1:]...)
				sellOrders = append(sellOrders[:j], sellOrders[j+1:]...)
				i-- // Adjust index after removal
				break
			}
		}
	}
}

// ProcessOrders using struct
func ProcessOrders(orderChan <-chan Order, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range orderChan {
		AddOrder(order)
		MatchOrders()
	}
}

func main() {
	orderChan := make(chan Order, 10)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go ProcessOrders(orderChan, &wg)
	}

	orders := []Order{
		{1, "buy", "AAPL", 100, 150.0},
		{2, "sell", "AAPL", 100, 150.0},
		{3, "buy", "GOOGL", 50, 2800.0},
		{4, "sell", "GOOGL", 50, 2800.0},
		{5, "buy", "TSLA", 200, 700.0},
		{6, "sell", "TSLA", 200, 700.0},
	}

	for _, order := range orders {
		orderChan <- order
		time.Sleep(100 * time.Millisecond)
	}

	close(orderChan)
	wg.Wait()

	fmt.Println("Final Order Book:")
	fmt.Println("Buy Orders:", buyOrders)
	fmt.Println("Sell Orders:", sellOrders)
}

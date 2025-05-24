package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Server represents a server to monitor
type Server struct {
	Name    string
	Address string
}

// HealthCheck simulates checking the health of a server
func HealthCheck(ctx context.Context, server Server) error {
	// Simulate network latency
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	// Simulate server health (randomly fail 10% of the time)
	if rand.Float32() < 0.1 {
		return fmt.Errorf("server %s is down", server.Name)
	}

	return nil
}

// MonitorServer monitors a single server and logs its health status
func MonitorServer(ctx context.Context, server Server, wg *sync.WaitGroup, results chan<- string, alerts chan<- string) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Stopping monitoring for server: %s\n", server.Name)
			return
		default:
			// Create a timeout context for the health check
			healthCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
			defer cancel()

			// Perform health check
			if err := HealthCheck(healthCtx, server); err != nil {
				results <- fmt.Sprintf("Server %s is DOWN: %v", server.Name, err)
				alerts <- fmt.Sprintf("ALERT: Server %s is DOWN!", server.Name)
			} else {
				results <- fmt.Sprintf("Server %s is UP", server.Name)
			}

			// Wait before the next health check
			time.Sleep(2 * time.Second)
		}
	}
}

// Logger logs results to a file or console
func Logger(results <-chan string) {
	for result := range results {
		log.Println(result)
	}
}

// AlertManager sends alerts if a server is down
func AlertManager(alerts <-chan string) {
	for alert := range alerts {
		log.Println(alert)
		// Simulate sending an alert (e.g., email, Slack, etc.)
		fmt.Println("Sending alert:", alert)
	}
}

func main() {
	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Simulate a list of servers to monitor
	servers := []Server{
		{Name: "Server1", Address: "192.168.1.1"},
		{Name: "Server2", Address: "192.168.1.2"},
		{Name: "Server3", Address: "192.168.1.3"},
		// Add more servers here...
	}

	// Channels for logging and alerts
	results := make(chan string, 100)
	alerts := make(chan string, 100)

	// Start the logger and alert manager
	go Logger(results)
	go AlertManager(alerts)

	// WaitGroup to wait for all monitoring goroutines to finish
	var wg sync.WaitGroup

	// Start monitoring each server
	for _, server := range servers {
		wg.Add(1)
		go MonitorServer(ctx, server, &wg, results, alerts)
	}

	// Simulate running the system for a while
	time.Sleep(10 * time.Second)

	// Stop monitoring (cancel all goroutines)
	cancel()

	// Wait for all goroutines to finish
	wg.Wait()

	// Close channels
	close(results)
	close(alerts)

	log.Println("Server monitoring stopped.")
}

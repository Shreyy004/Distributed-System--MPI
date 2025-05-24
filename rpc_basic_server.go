package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Define a struct
type Arithmetic struct{}

// Define an RPC method
func (a *Arithmetic) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Define arguments struct
type Args struct {
	A, B int
}

func main() {
	// Register the service
	arith := new(Arithmetic)
	rpc.Register(arith)

	// Start listening for connections
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("RPC Server listening on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}

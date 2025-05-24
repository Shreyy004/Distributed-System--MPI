package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Calculator struct{}

func (c *Calculator) Sub(args *Args, res *int) error {
	*res = args.A - args.B
	return nil

}

func (c *Calculator) Add(args *Args, res *int) error {
	*res = args.A + args.B
	return nil
}

type Args struct {
	A, B int
}

func main() {
	calc := new(Calculator)
	rpc.Register(calc)

	lister, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	defer lister.Close()
	fmt.Print("server connect to port 1234...")

	for {
		conn, err := lister.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}

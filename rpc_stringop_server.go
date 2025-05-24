package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Request struct for single string operations
type StringRequest struct {
	Str     string
	Char    string
	Pos     int
}

// Request struct for concatenation
type ConcatRequest struct {
	Str1 string
	Str2 string
}

// Response struct for returning string results
type Response struct {
	Result string
}

// StringService struct
type StringService struct{}

// Insert a character into a string at the given position
func (s *StringService) Insert(req StringRequest, res *Response) error {
	if req.Pos < 0 || req.Pos > len(req.Str) {
		res.Result = "Invalid position!"
		return nil
	}
	res.Result = req.Str[:req.Pos] + req.Char + req.Str[req.Pos:]
	return nil
}

// Delete a character from a string at the given position
func (s *StringService) Delete(req StringRequest, res *Response) error {
	if req.Pos < 0 || req.Pos >= len(req.Str) {
		res.Result = "Invalid position!"
		return nil
	}
	res.Result = req.Str[:req.Pos] + req.Str[req.Pos+1:]
	return nil
}

// Modify a character in a string at the given position
func (s *StringService) Modify(req StringRequest, res *Response) error {
	if req.Pos < 0 || req.Pos >= len(req.Str) {
		res.Result = "Invalid position!"
		return nil
	}
	res.Result = req.Str[:req.Pos] + req.Char + req.Str[req.Pos+1:]
	return nil
}

// Concatenate two strings
func (s *StringService) Concatenate(req ConcatRequest, res *Response) error {
	res.Result = req.Str1 + req.Str2
	return nil
}

func main() {
	service := new(StringService)
	rpc.Register(service)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("RPC String Service started on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

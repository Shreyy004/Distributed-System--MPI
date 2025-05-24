package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

// StringService provides methods for string manipulation
type StringService struct{}

// Request structure
type StringRequest struct {
	Str     string
	Char    string
	Index   int
	Str2    string
}

// Response structure
type StringResponse struct {
	Result string
}

// Insert a character into a string
func (s *StringService) Insert(req StringRequest, res *StringResponse) error {
	if req.Index < 0 || req.Index > len(req.Str) {
		return errors.New("invalid index")
	}
	res.Result = req.Str[:req.Index] + req.Char + req.Str[req.Index:]
	return nil
}

// Delete a character from a string
func (s *StringService) Delete(req StringRequest, res *StringResponse) error {
	if req.Index < 0 || req.Index >= len(req.Str) {
		return errors.New("invalid index")
	}
	res.Result = req.Str[:req.Index] + req.Str[req.Index+1:]
	return nil
}

// Modify a character in a string
func (s *StringService) Modify(req StringRequest, res *StringResponse) error {
	if req.Index < 0 || req.Index >= len(req.Str) {
		return errors.New("invalid index")
	}
	res.Result = req.Str[:req.Index] + req.Char + req.Str[req.Index+1:]
	return nil
}

// Concatenate two strings
func (s *StringService) Concat(req StringRequest, res *StringResponse) error {
	res.Result = req.Str + req.Str2
	return nil
}

func main() {
	strService := new(StringService)
	rpc.Register(strService)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is running on port 1234...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

package main

import (
	"fmt"
	"net/rpc"
)

// Request struct for string operations
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

// Response struct for results
type Response struct {
	Result string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	for {
		fmt.Println("\nString Operations - Select an option:")
		fmt.Println("1. Insert a character")
		fmt.Println("2. Delete a character")
		fmt.Println("3. Modify a character")
		fmt.Println("4. Concatenate two strings")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var str, char string
			var pos int
			fmt.Print("Enter the string: ")
			fmt.Scanln(&str)
			fmt.Print("Enter the character to insert: ")
			fmt.Scanln(&char)
			fmt.Print("Enter position: ")
			fmt.Scanln(&pos)

			req := StringRequest{Str: str, Char: char, Pos: pos}
			var res Response
			err = client.Call("StringService.Insert", req, &res)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Result:", res.Result)
			}

		case 2:
			var str string
			var pos int
			fmt.Print("Enter the string: ")
			fmt.Scanln(&str)
			fmt.Print("Enter position to delete: ")
			fmt.Scanln(&pos)

			req := StringRequest{Str: str, Pos: pos}
			var res Response
			err = client.Call("StringService.Delete", req, &res)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Result:", res.Result)
			}

		case 3:
			var str, char string
			var pos int
			fmt.Print("Enter the string: ")
			fmt.Scanln(&str)
			fmt.Print("Enter the new character: ")
			fmt.Scanln(&char)
			fmt.Print("Enter position to modify: ")
			fmt.Scanln(&pos)

			req := StringRequest{Str: str, Char: char, Pos: pos}
			var res Response
			err = client.Call("StringService.Modify", req, &res)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Result:", res.Result)
			}

		case 4:
			var str1, str2 string
			fmt.Print("Enter first string: ")
			fmt.Scanln(&str1)
			fmt.Print("Enter second string: ")
			fmt.Scanln(&str2)

			req := ConcatRequest{Str1: str1, Str2: str2}
			var res Response
			err = client.Call("StringService.Concatenate", req, &res)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Result:", res.Result)
			}

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

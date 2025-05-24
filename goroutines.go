package main
import (
"fmt"
"time"
)

func take_input()(string, string, int) {
var first_name string
var last_name string
var num_tickets int
fmt.Println("enter first name: ")
fmt.Scanf("%s", &first_name)
fmt.Println("enter last name: ")
fmt.Scanf("%s", &last_name)
fmt.Println("enter number of tickets: ")
fmt.Scanf("%d", &num_tickets)
return first_name, last_name, num_tickets
}

func display_details(abc string, def int) {
fmt.Printf("%v booked %v tickets\n", abc, def)

}

func greetings() {
fmt.Println("Welcome!")
}

func main() {
var f_name, l_name, full_name string
var tickets int
f_name, l_name, tickets = take_input()
full_name = f_name + " " + l_name
go greetings()
time.Sleep(1*time.Second)
display_details(full_name, tickets)
}

package main
import "fmt"

type details struct {
name string
tickets int
}

func main() {
const conferenceName string = "go conference"
fmt.Printf("welcome to %v\n", conferenceName)

const conferenceTicket int = 50
var remainingTicket int = 50

fmt.Printf("we have a total of %v tickets, and %v tickets are still available\n", conferenceTicket, remainingTicket) 
fmt.Println("enter details: ")

var bookingDetails[] details
var first_name string
var last_name string
var full_name string
var num_tickets int

for remainingTicket>0 {

fmt.Println("enter first name: ")
fmt.Scanf("%s", &first_name)

fmt.Println("enter last name: ")
fmt.Scanf("%s", &last_name)
full_name = first_name + " " + last_name

fmt.Println("enter number of tickets: ")
fmt.Scanf("%d", &num_tickets)

if remainingTicket - num_tickets < 0 {
fmt.Printf("limited tickets available. %v tickets are still available\n", remainingTicket)
} else {
remainingTicket = remainingTicket - num_tickets
bookingDetails = append(bookingDetails, details{full_name, num_tickets})
fmt.Printf("%v successfully booked %v tickets\n", full_name, num_tickets)
}


}

}

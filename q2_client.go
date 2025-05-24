package main
import (
"fmt"
"net/rpc"
)

func main() {
for {
fmt.Println("\n--- Hotel Booking System ---")
fmt.Println("1. Book a room")
fmt.Println("2. Cancel a booking")
fmt.Println("3. Check room availability")
fmt.Println("4. Exit")
fmt.Print("Choose an option: ")
var choice, roomType int
fmt.Scanf("%d %d", &choice, &roomType)
if (choice == 4) {
break
}

}
